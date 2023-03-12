package repository

import (
	"project-intern-bcc/src/business/entity"
	"project-intern-bcc/src/lib/midtrans"

	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
)

type PremiumOrderRepository interface {
	Create(premiumOrder entity.PremiumOrders)(entity.PremiumOrders,*coreapi.ChargeResponse,error)
	FindByOrderCode(orderCode string)(entity.PremiumOrders,error)
	Update(order entity.PremiumOrders)(error)
}

type premiumOrderRepository struct {
	db *gorm.DB
	midtrans midtrans.MidtransInterface
}

func NewPremiumOrderRepository(db *gorm.DB, midtrans midtrans.MidtransInterface) PremiumOrderRepository {
	return &premiumOrderRepository{
		db:db,
		midtrans: midtrans,
	}
}

func (h *premiumOrderRepository) Create(premiumOrder entity.PremiumOrders)(entity.PremiumOrders,*coreapi.ChargeResponse,error){
	err:=h.db.Create(&premiumOrder).Error
	if err!=nil{
		return premiumOrder,nil,err
	}

	resp,err:=h.midtrans.CreatePremiumTransaction(premiumOrder)
	return premiumOrder,resp,err
}

func (h *premiumOrderRepository) FindByOrderCode(orderCode string)(entity.PremiumOrders,error){
	var order entity.PremiumOrders
	err:=h.db.First(&order, "order_code=?",orderCode).Error
	return order,err
}

func (h *premiumOrderRepository) Update(order entity.PremiumOrders)(error){
	err:=h.db.Save(&order).Error
	return err
}