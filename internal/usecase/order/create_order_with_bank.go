package order

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	d "github.com/phincon-backend/laza/domain/repositories/cart"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/helper"
	"strings"
	"time"
)

type CreateOrderWithBankUsecase struct {
	insertOrder              repositories.InsertAction[model.Order]
	getAddressById           repositories.GetByIdAction[model.Address]
	chargeMidtrans           midtranscore.ChargeMidtransAction
	getOrder                 repositories.GetByIdAction[model.Order]
	getProduct               repositories.GetByIdAction[model.Product]
	insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail]
	getCategory              repositories.GetByIdAction[model.Category]
	getBrand                 repositories.GetByIdAction[model.Brand]
	insertPaymentMethod      repositories.InsertAction[model.PaymentMethod]
	getCartByIdRepo          d.GetCartByIdAction
}

func NewCreateOrderWithBankUsecase(
	insertOrder repositories.InsertAction[model.Order],
	getAddressById repositories.GetByIdAction[model.Address],
	chargeMidtrans midtranscore.ChargeMidtransAction,
	getOrder repositories.GetByIdAction[model.Order],
	getProduct repositories.GetByIdAction[model.Product],
	insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail],
	getCategory repositories.GetByIdAction[model.Category],
	getBrand repositories.GetByIdAction[model.Brand],
	insertPaymentMethod repositories.InsertAction[model.PaymentMethod],
	getCartByIdRepo d.GetCartByIdAction,
) *CreateOrderWithBankUsecase {
	return &CreateOrderWithBankUsecase{
		insertOrder:              insertOrder,
		getAddressById:           getAddressById,
		chargeMidtrans:           chargeMidtrans,
		getOrder:                 getOrder,
		getProduct:               getProduct,
		insertProductOrderDetail: insertProductOrderDetail,
		getCategory:              getCategory,
		getBrand:                 getBrand,
		insertPaymentMethod:      insertPaymentMethod,
		getCartByIdRepo:          getCartByIdRepo,
	}
}

func (uc *CreateOrderWithBankUsecase) Execute(userId uint64, addressId int, bank string) (*model.Order, *model.PaymentMethod, error) {
	// check if address exists
	address, err := uc.getAddressById.GetById(addressId)
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
	var grossAmount int = 0
	productsDetails := make([]model.ProductOrderDetail, 0)
	productCarts, err := uc.getCartByIdRepo.GetCartById(userId)
	for _, productCart := range productCarts {
		productTemp, err := uc.getProduct.GetById(productCart.Id)
		if err != nil {
			return nil, nil, err
		}

		categoryTemp, err := uc.getCategory.GetById(productTemp.CategoryId)
		if err != nil {
			return nil, nil, err
		}

		brandTemp, err := uc.getBrand.GetById(productTemp.BrandId)
		if err != nil {
			return nil, nil, err
		}

		productsDetails = append(productsDetails,
			model.ProductOrderDetail{
				Name:        productTemp.Name,
				Description: productTemp.Description,
				ImageUrl:    productTemp.ImageUrl,
				Price:       int(productTemp.Price),
				Category:    categoryTemp.Category,
				BrandName:   brandTemp.Name,
				Quantity:    productCart.Quantity,
				Size:        "",
				TotalPrice:  int(productTemp.Price) * productCart.Quantity,
				OrderId:     orderNumber,
			},
		)
	}

	// charge bank to midtrans
	var paymentReq coreapi.ChargeReq
	if strings.ToLower(bank) == "mandiri" {
		paymentReq = coreapi.ChargeReq{
			PaymentType: "echannel",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  orderNumber,
				GrossAmt: int64(grossAmount),
			},
			EChannel: &coreapi.EChannelDetail{
				BillInfo1: "Payment ",
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
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: midtrans.Bank(bank),
			},
		}
	}
	respondMd, err := uc.chargeMidtrans.ChargeMidtrans(&paymentReq)
	if err != nil {
		return nil, nil, err
	}

	// parsing time
	parsedTime, err := time.Parse(respondMd.ExpiryTime, respondMd.ExpiryTime)
	if err != nil {
		return nil, nil, err
	}

	// insert payment method to db
	var paymentMethod model.PaymentMethod
	if strings.ToLower(bank) == "mandiri" {
		paymentMethod = model.PaymentMethod{
			PaymentMethod: "bank",
			Bank:          bank,
			BillerCode:    respondMd.BillerCode,
			BillKey:       respondMd.BillKey,
			ExpiryTime:    parsedTime,
		}
	} else if strings.ToLower(bank) == "permata" {
		paymentMethod = model.PaymentMethod{
			PaymentMethod: "bank",
			Bank:          bank,
			VANumber:      respondMd.PermataVaNumber,
			ExpiryTime:    parsedTime,
		}
	} else {
		paymentMethod = model.PaymentMethod{
			PaymentMethod: "bank",
			Bank:          bank,
			VANumber:      respondMd.VaNumbers[0].VANumber,
			ExpiryTime:    parsedTime,
		}
	}
	paymentMethod, err = uc.insertPaymentMethod.Insert(paymentMethod)
	if err != nil {
		return nil, nil, err
	}

	// insert order to db
	order := model.Order{
		Id:              orderNumber,
		Amount:          int64(grossAmount),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		PaidAt:          time.Time{},
		ExpiryDate:      paymentMethod.ExpiryTime,
		ShippingFee:     helper.GenerateShippingFee(address),
		AdminFee:        helper.GenerateAdminFee(),
		OrderStatus:     respondMd.TransactionStatus,
		UserId:          userId,
		AddressId:       uint64(addressId),
		PaymentMethodId: paymentMethod.Id,
	}
	orderRespond, err := uc.insertOrder.Insert(order)
	if err != nil {
		return nil, nil, err
	}

	for _, productsDetail := range productsDetails {
		_, err = uc.insertProductOrderDetail.Insert(productsDetail)
		if err != nil {
			return nil, nil, err
		}
	}

	return &orderRespond, &paymentMethod, nil
}
