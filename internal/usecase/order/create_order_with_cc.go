package order

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
	"time"
)

type CreateOrderWithCCUsecase struct {
	insertOrder        repositories.InsertAction[model.Order]
	getAddressById     repositories.GetByIdAction[model.Address]
	chargeMidtrans     midtranscore.ChargeMidtransAction
	getCCToken         midtranscore.FetchMidtransCCTokenAction
	insertCreditCard   repositories.InsertAction[model.CreditCard]
	getOrder           repositories.GetByIdAction[model.Order]
	getProduct         repositories.GetByIdAction[model.Product]
	insertProductOrder repositories.InsertAction[model.ProductOrder]
}

func NewCreateOrderWithCCUsecase(
	insertOrder repositories.InsertAction[model.Order],
	getAddressById repositories.GetByIdAction[model.Address],
	chargeMidtrans midtranscore.ChargeMidtransAction,
	getCCToken midtranscore.FetchMidtransCCTokenAction,
	insertCreditCard repositories.InsertAction[model.CreditCard],
	getOrder repositories.GetByIdAction[model.Order],
	getProduct repositories.GetByIdAction[model.Product],
	insertProductOrder repositories.InsertAction[model.ProductOrder],
) *CreateOrderWithCCUsecase {
	return &CreateOrderWithCCUsecase{
		insertOrder:        insertOrder,
		getAddressById:     getAddressById,
		chargeMidtrans:     chargeMidtrans,
		getCCToken:         getCCToken,
		insertCreditCard:   insertCreditCard,
		getOrder:           getOrder,
		getProduct:         getProduct,
		insertProductOrder: insertProductOrder,
	}
}

func (uc *CreateOrderWithCCUsecase) Execute(userId uint64, addressId int, cc model.CreditCard, cvv string, products []requests.ProductOrder) (*model.Order, *model.CreditCard, error) {
	// check if address exists
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

	// get cc token
	cardTokenResponse, errMd := uc.getCCToken.FetchMidtransCCToken(cc.CardNumber, cc.ExpiredMonth, cc.ExpiredYear, cvv)
	byteArr, _ := json.Marshal(errMd)
	fmt.Println("error get token: ", errMd)
	fmt.Println("cc token: ", cardTokenResponse.TokenID)
	if errMd != nil {
		fmt.Println("masuk if error", errMd)
		return nil, nil, errMd.RawError
	}

	// marshalling the structure
	byteArr, _ = json.Marshal(cardTokenResponse)
	fmt.Println("respond Card Token Response", string(byteArr))

	// charge cc to midtrans
	var paymentReq coreapi.ChargeReq
	paymentReq = coreapi.ChargeReq{
		PaymentType: "credit_card",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderNumber,
			GrossAmt: int64(grossAmount),
		},
		CreditCard: &coreapi.CreditCardDetails{
			TokenID:        cardTokenResponse.TokenID,
			Authentication: false,
			Bank:           cardTokenResponse.Bank,
			CallbackType:   cardTokenResponse.RedirectURL,
		},
	}
	responseMd, err := uc.chargeMidtrans.ChargeMidtrans(&paymentReq)
	// marshalling the structure
	byteArr, _ = json.Marshal(responseMd)
	fmt.Println("respond midtrans charge", string(byteArr))
	if err != nil {
		fmt.Println("error midtrans charge: ", err)
		return nil, nil, err
	}

	fmt.Println("creditcard : ", cc)

	// insert order to db
	order := model.Order{
		Id:          orderNumber,
		Amount:      int64(grossAmount),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserId:      userId,
		OrderStatus: responseMd.TransactionStatus,
		AddressId:   uint64(addressId),
		CreditCardId: sql.NullInt64{
			Int64: int64(cc.Id),
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

	return &orderRespond, &cc, nil
}
