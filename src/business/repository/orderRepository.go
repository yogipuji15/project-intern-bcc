package repository

import (
	"mime/multipart"
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/lib/midtrans"
	"project-intern-bcc/src/lib/storage"

	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(speaker entity.Speakers,user entity.UserResponse,order entity.Orders,rundown *multipart.FileHeader,script *multipart.FileHeader)(entity.Orders,*coreapi.ChargeResponse,error)
	FindByOrderCode(orderCode string)(entity.Orders,error)
	Update(order entity.Orders)(error)
	FindAll(userId any, pagination entity.Pagination)([]entity.Orders,*entity.Pagination,error)
	CreateMidtransTransaction(order entity.Orders,speaker entity.Speakers)(*coreapi.ChargeResponse,error)
}

type orderRepository struct {
	db *gorm.DB
	storage storage.StorageInterface
	midtrans midtrans.MidtransInterface
}

func NewOrderRepository(db *gorm.DB, storage storage.StorageInterface, midtrans midtrans.MidtransInterface) OrderRepository {
	return &orderRepository{
		db: db,
		storage: storage,
		midtrans: midtrans,
	}
}

func (h *orderRepository) Create(speaker entity.Speakers,user entity.UserResponse,order entity.Orders,rundown *multipart.FileHeader,script *multipart.FileHeader)(entity.Orders,*coreapi.ChargeResponse,error){
	rundownLink, err :=h.storage.UploadFile(rundown)
	if err != nil{
		return order,nil,err
	}
	
	scriptLink, err :=h.storage.UploadFile(script)
	if err != nil{
		return order,nil,err
	}

	order.Rundown=rundownLink
	order.Script=scriptLink
	
	err=h.db.Create(&order).Error
	
	if user.Role=="premium-user"{
		resp,err:=h.midtrans.CreateTransaction(order,speaker)
		if err!=nil{
			return order,nil,err
		}
		return order,resp,err
	}

	return order,nil,err
}

func (h *orderRepository) Update(order entity.Orders)(error){
	err:=h.db.Save(&order).Error
	return err
}

func (h *orderRepository) FindByOrderCode(orderCode string)(entity.Orders,error){
	var order entity.Orders
	err:=h.db.First(&order, "order_code=?",orderCode).Error
	return order,err
}

func (h *orderRepository) FindAll(userId any, pagination entity.Pagination)([]entity.Orders,*entity.Pagination,error){
	pg:= entity.FormatPaginationParam(pagination)

	var orders []entity.Orders
	err:= h.db.Preload("Speaker").Where("user_id = ?",userId).Offset(int(pg.Offset)).Limit(int(pg.Limit)).Find(&orders).Error
	if err!=nil{
		return nil,nil,err
	}

	err = h.db.Model(&orders).Preload("Speaker").Where("user_id = ?",userId).Count(&pg.TotalElement).Error
	if err!=nil{
		return nil,nil,err
	}

	pg.ProcessPagination(int64(len(orders)))

	return orders,&pg,err
}

func (h *orderRepository)  CreateMidtransTransaction(order entity.Orders,speaker entity.Speakers)(*coreapi.ChargeResponse,error){
	resp,err:=h.midtrans.CreateTransaction(order,speaker)
	return resp,err
}