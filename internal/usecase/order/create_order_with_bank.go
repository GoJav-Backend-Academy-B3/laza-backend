package order

import (
	"database/sql"
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/helper"
	"strings"
	"time"
)

type CreateOrderWithBankUsecase struct {
	insertOrder           repositories.InsertAction[model.Order]
	getAddressById        repositories.GetByIdAction[model.Address]
	chargeMidtrans        midtranscore.ChargeMidtransAction
	insertTransactionBank repositories.InsertAction[model.TransactionBank]
	getOrder              repositories.GetByIdAction[model.Order]
	getProduct            repositories.GetByIdAction[model.Product]
	insertProductOrder    repositories.InsertAction[model.ProductOrder]
}

func NewCreateOrderWithBankUsecase(
	insertOrder repositories.InsertAction[model.Order],
	getAddressById repositories.GetByIdAction[model.Address],
	chargeMidtrans midtranscore.ChargeMidtransAction,
	insertTransactionBank repositories.InsertAction[model.TransactionBank],
	getOrder repositories.GetByIdAction[model.Order],
	getProduct repositories.GetByIdAction[model.Product],
	insertProductOrder repositories.InsertAction[model.ProductOrder],
) *CreateOrderWithBankUsecase {
	return &CreateOrderWithBankUsecase{
		insertOrder:           insertOrder,
		getAddressById:        getAddressById,
		chargeMidtrans:        chargeMidtrans,
		insertTransactionBank: insertTransactionBank,
		getOrder:              getOrder,
		getProduct:            getProduct,
		insertProductOrder:    insertProductOrder,
	}
}

func (uc *CreateOrderWithBankUsecase) Execute(userId uint64, addressId int, bank string, products []request.ProductOrder) (*model.Order, *model.TransactionBank, error) {
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

	// charge bank to midtrans
	var paymentReq coreapi.ChargeReq
	if strings.ToLower(bank) == "mandiri" {
		paymentReq = coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderNumber,
				GrossAmt: int64(grossAmount),
			},
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "Payment For:",
				BillInfo2: "Laza with order ID: " + orderNumber,
			},
		}
	} else {
		paymentReq = coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderNumber,
				GrossAmt: int64(grossAmount),
			},
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "Payment For:",
				BillInfo2: "Laza with order ID: " + orderNumber,
			},
		}
	}
	RespondMd, err := uc.chargeMidtrans.ChargeMidtrans(&paymentReq)
	fmt.Println("Bank Respond: ", RespondMd)
	if err != nil {
		return nil, nil, err
	}

	// insert bank to db
	var transactionBankModel model.TransactionBank
	if strings.ToLower(bank) == "mandiri" {
		transactionBankModel = model.TransactionBank{
			BankCode:   RespondMd.Bank,
			BillerCode: "",
			VANumber:   RespondMd.VaNumbers[0].VANumber,
		}
	} else {
		transactionBankModel = model.TransactionBank{
			BankCode:   RespondMd.Bank,
			BillerCode: RespondMd.BillerCode,
			VANumber:   RespondMd.BillKey,
		}
	}
	insertRes, err := uc.insertTransactionBank.Insert(transactionBankModel)
	fmt.Println("insert result:", insertRes)
	if err != nil {
		return nil, nil, err
	}

	// insert order to db
	order := model.Order{
		Id:            orderNumber,
		Amount:        int64(grossAmount),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		UserId:        userId,
		OrderStatusId: 1,
		AddressId:     uint64(addressId),
		TransactionBankId: sql.NullInt64{
			Int64: int64(insertRes.Id),
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

	return &orderRespond, &transactionBankModel, nil
}
