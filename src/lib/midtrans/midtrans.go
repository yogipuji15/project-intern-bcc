package midtrans

import (
	"os"
	"project-intern-bcc/src/business/entity"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)



type MidtransInterface interface {
	CreateTransaction(order entity.Orders,speaker entity.Speakers)(*coreapi.ChargeResponse,error)
}

type midtransStruct struct {
	midtransClient coreapi.Client
}
func Init() MidtransInterface {

	return &midtransStruct{}
}

func (m *midtransStruct) CreateTransaction(order entity.Orders,speaker entity.Speakers)(*coreapi.ChargeResponse,error) {
	m.midtransClient = coreapi.Client{}
	m.midtransClient.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	
	var req *coreapi.ChargeReq
	if order.PaymentID==1||order.PaymentID==2{
		req = &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  order.OrderCode,
				GrossAmt: int64(order.TotalPrice),
			},
			Items: &[]midtrans.ItemDetails{
				midtrans.ItemDetails{
					ID		: strconv.Itoa(int(speaker.ID)),
					Price	: int64(speaker.Price),
					Qty		: int32(order.Duration),
					Name	: speaker.Name,
				},
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: "bca",
			},
		}
		
	}else if order.PaymentID==3{
		req = &coreapi.ChargeReq{
			PaymentType: "gopay",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  order.OrderCode,
				GrossAmt: int64(order.TotalPrice),
				
			},
			Items: &[]midtrans.ItemDetails{
				midtrans.ItemDetails{
					ID		: strconv.Itoa(int(speaker.ID)),
					Price	: int64(speaker.Price),
					Qty		: int32(order.Duration),
					Name	: speaker.Name,
				},
			},
		}
	}
	
	
	resp, _ := m.midtransClient.ChargeTransaction(req)
	return resp,nil
	
}

