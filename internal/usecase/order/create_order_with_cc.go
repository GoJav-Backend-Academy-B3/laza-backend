package order

import (
	"database/sql"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	d "github.com/phincon-backend/laza/domain/repositories/cart"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/helper"
	"time"
)

type CreateOrderWithCCUsecase struct {
	insertOrder              repositories.InsertAction[model.Order]
	getAddressById           repositories.GetByIdAction[model.Address]
	chargeMidtrans           midtranscore.ChargeMidtransAction
	getCCToken               midtranscore.FetchMidtransCCTokenAction
	insertCreditCard         repositories.InsertAction[model.CreditCard]
	getOrder                 repositories.GetByIdAction[model.Order]
	getProduct               repositories.GetByIdAction[model.Product]
	insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail]
	getCategory              repositories.GetByIdAction[model.Category]
	getBrand                 repositories.GetByIdAction[model.Brand]
	insertPaymentMethod      repositories.InsertAction[model.PaymentMethod]
	getCartByIdRepo          d.GetCartByIdAction
}

func NewCreateOrderWithCCUsecase(
	insertOrder repositories.InsertAction[model.Order],
	getAddressById repositories.GetByIdAction[model.Address],
	chargeMidtrans midtranscore.ChargeMidtransAction,
	getCCToken midtranscore.FetchMidtransCCTokenAction,
	insertCreditCard repositories.InsertAction[model.CreditCard],
	getOrder repositories.GetByIdAction[model.Order],
	getProduct repositories.GetByIdAction[model.Product],
	insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail],
	getCategory repositories.GetByIdAction[model.Category],
	getBrand repositories.GetByIdAction[model.Brand],
	insertPaymentMethod repositories.InsertAction[model.PaymentMethod],
	getCartByIdRepo d.GetCartByIdAction,
) *CreateOrderWithCCUsecase {
	return &CreateOrderWithCCUsecase{
		insertOrder:              insertOrder,
		getAddressById:           getAddressById,
		chargeMidtrans:           chargeMidtrans,
		getCCToken:               getCCToken,
		insertCreditCard:         insertCreditCard,
		getOrder:                 getOrder,
		getProduct:               getProduct,
		insertProductOrderDetail: insertProductOrderDetail,
		getCategory:              getCategory,
		getBrand:                 getBrand,
		insertPaymentMethod:      insertPaymentMethod,
		getCartByIdRepo:          getCartByIdRepo,
	}
}

func (uc *CreateOrderWithCCUsecase) Execute(userId uint64, addressId int, cc model.CreditCard, cvv string) (*model.Order, *model.PaymentMethod, error) {
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
		grossAmount += int(productTemp.Price) * productCart.Quantity
	}

	// get cc token
	cardTokenResponse, errMd := uc.getCCToken.FetchMidtransCCToken(cc.CardNumber, cc.ExpiredMonth, cc.ExpiredYear, cvv)
	if errMd != nil {
		return nil, nil, errMd.RawError
	}

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
	if err != nil {
		return nil, nil, err
	}

	// parsing time
	parsedTime, err := time.Parse(responseMd.ExpiryTime, responseMd.ExpiryTime)
	if err != nil {
		return nil, nil, err
	}

	// insert payment method to db
	paymentMethod, err := uc.insertPaymentMethod.Insert(model.PaymentMethod{
		Id:               0,
		PaymentMethod:    "credit_card",
		CreditCardNumber: cc.CardNumber,
		RedirectUrl:      responseMd.RedirectURL,
		ExpiryTime:       parsedTime,
	})
	if err != nil {
		return nil, nil, err
	}

	// insert order to db
	order := model.Order{
		Id:              orderNumber,
		Amount:          int64(grossAmount),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		PaidAt:          sql.NullTime{Valid: false},
		ExpiryDate:      paymentMethod.ExpiryTime,
		ShippingFee:     helper.GenerateShippingFee(address),
		AdminFee:        helper.GenerateAdminFee(),
		OrderStatus:     responseMd.TransactionStatus,
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
