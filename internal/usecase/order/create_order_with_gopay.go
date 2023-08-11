package order

import (
	"database/sql"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"time"
)

type CreateOrderWithGopayUsecase struct {
	insertOrder        repositories.InsertAction[model.Order]
	getAddressById     repositories.GetByIdAction[model.Address]
	chargeGopay        midtranscore.ChargeMidtransAction
	insertGopay        repositories.InsertAction[model.Gopay]
	getOrder           repositories.GetByIdAction[model.Order]
	getProduct         repositories.GetByIdAction[model.Product]
	insertProductOrder repositories.InsertAction[model.ProductOrder]
}

func NewCreateOrderWithGopayUsecase(
	insertOrder repositories.InsertAction[model.Order],
	getAddressById repositories.GetByIdAction[model.Address],
	chargeGopay midtranscore.ChargeMidtransAction,
	insertGopay repositories.InsertAction[model.Gopay],
	getOrder repositories.GetByIdAction[model.Order],
	getProduct repositories.GetByIdAction[model.Product],
	insertProductOrder repositories.InsertAction[model.ProductOrder],
) *CreateOrderWithGopayUsecase {
	return &CreateOrderWithGopayUsecase{
		insertOrder:        insertOrder,
		getAddressById:     getAddressById,
		chargeGopay:        chargeGopay,
		insertGopay:        insertGopay,
		getOrder:           getOrder,
		getProduct:         getProduct,
		insertProductOrder: insertProductOrder,
	}
}

func (uc *CreateOrderWithGopayUsecase) Execute(userId uint64, addressId int, callbackUrl string, products []requests.ProductOrder) (*model.Order, *model.Gopay, error) {
	// Check if address exists
	_, err := uc.getAddressById.GetById(addressId)
	if err != nil {
		return nil, nil, err
	}

	// Generate order number
	var orderNumber string
	for true {
		orderNumber = helper.GenerateOrderNumber()
		_, err := uc.getOrder.GetById(orderNumber)
		if err != nil {
			break
		}
	}

	// count gross amount
	var grossAmount float64 = 0
	var productOrders = make([]model.ProductOrder, 0)
	for _, product := range products {
		productTemp, err := uc.getProduct.GetById(product.Id)
		if err != nil {
			return nil, nil, err
		}

		productOrderTemp := model.ProductOrder{
			ProductId: productTemp.Id,
			OrderId:   orderNumber,
			Quantity:  uint16(product.Quantity),
			Price:     productTemp.Price * float64(product.Quantity),
		}

		productOrders = append(productOrders, productOrderTemp)

		grossAmount += productTemp.Price * float64(product.Quantity)
	}

	// Charge gopay to midtrans
	paymentReq := coreapi.ChargeReq{
		PaymentType: "gopay",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderNumber,
			GrossAmt: int64(grossAmount),
		},
	}
	gopayRespondMd, err := uc.chargeGopay.ChargeMidtrans(&paymentReq)
	if err != nil {
		return nil, nil, err
	}

	// insert gopay to db
	gopayRespond, err := uc.insertGopay.Insert(model.Gopay{
		QRCode:        gopayRespondMd.Actions[0].URL,
		Deeplink:      gopayRespondMd.Actions[1].URL,
		GetStatusLink: gopayRespondMd.Actions[2].URL,
		CancelLink:    gopayRespondMd.Actions[3].URL,
		ExpiryTime:    time.Now(),
	})
	if err != nil {
		return nil, nil, err
	}

	// insert order to db
	order := model.Order{
		Id:          orderNumber,
		Amount:      int64(grossAmount),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserId:      userId,
		OrderStatus: gopayRespondMd.TransactionStatus,
		AddressId:   uint64(addressId),
		GopayId: sql.NullInt64{
			Int64: int64(gopayRespond.Id),
			Valid: true,
		},
	}
	orderRespond, err := uc.insertOrder.Insert(order)
	if err != nil {
		return nil, nil, err
	}

	for _, productOrder := range productOrders {
		_, err = uc.insertProductOrder.Insert(productOrder)
		if err != nil {
			return nil, nil, err
		}
	}

	return &orderRespond, &gopayRespond, nil
}
