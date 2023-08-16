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

type CreateOrderWithGopayUsecase struct {
	insertOrder              repositories.InsertAction[model.Order]
	getAddressById           repositories.GetByIdAction[model.Address]
	chargeGopay              midtranscore.ChargeMidtransAction
	getOrder                 repositories.GetByIdAction[model.Order]
	getProduct               repositories.GetByIdAction[model.Product]
	insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail]
	getCategory              repositories.GetByIdAction[model.Category]
	getBrand                 repositories.GetByIdAction[model.Brand]
	insertPaymentMethod      repositories.InsertAction[model.PaymentMethod]
	getCartByIdRepo          d.GetCartByIdAction
	deleteCartByUser         repositories.DeleteAction[model.Cart]
}

func NewCreateOrderWithGopayUsecase(insertOrder repositories.InsertAction[model.Order], getAddressById repositories.GetByIdAction[model.Address], chargeGopay midtranscore.ChargeMidtransAction, getOrder repositories.GetByIdAction[model.Order], getProduct repositories.GetByIdAction[model.Product], insertProductOrderDetail repositories.InsertAction[model.ProductOrderDetail], getCategory repositories.GetByIdAction[model.Category], getBrand repositories.GetByIdAction[model.Brand], insertPaymentMethod repositories.InsertAction[model.PaymentMethod], getCartByIdRepo d.GetCartByIdAction, deleteCartByUser repositories.DeleteAction[model.Cart]) *CreateOrderWithGopayUsecase {
	return &CreateOrderWithGopayUsecase{insertOrder: insertOrder, getAddressById: getAddressById, chargeGopay: chargeGopay, getOrder: getOrder, getProduct: getProduct, insertProductOrderDetail: insertProductOrderDetail, getCategory: getCategory, getBrand: getBrand, insertPaymentMethod: insertPaymentMethod, getCartByIdRepo: getCartByIdRepo, deleteCartByUser: deleteCartByUser}
}

func (uc *CreateOrderWithGopayUsecase) Execute(userId uint64, addressId int, callbackUrl string) (*model.Order, *model.PaymentMethod, error) {
	// Check if address exists
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
				Size:        productCart.Size,
				TotalPrice:  int(productTemp.Price) * productCart.Quantity,
				OrderId:     orderNumber,
			},
		)
		grossAmount += int(productTemp.Price) * productCart.Quantity
	}

	// Charge gopay to midtrans
	paymentReq := coreapi.ChargeReq{
		PaymentType: "gopay",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderNumber,
			GrossAmt: int64(grossAmount),
		},
		Gopay: &coreapi.GopayDetails{
			EnableCallback: true,
			CallbackUrl:    callbackUrl,
		},
	}
	gopayRespondMd, err := uc.chargeGopay.ChargeMidtrans(&paymentReq)
	if err != nil {
		return nil, nil, err
	}

	// parsing time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", gopayRespondMd.ExpiryTime)
	if err != nil {
		return nil, nil, err
	}

	// insert payment method to db
	paymentMethod, err := uc.insertPaymentMethod.Insert(model.PaymentMethod{
		PaymentMethod: "gopay",
		QRCodeUrl:     gopayRespondMd.Actions[0].URL,
		Deeplink:      gopayRespondMd.Actions[1].URL,
		ExpiryTime:    parsedTime,
	})
	if err != nil {
		return nil, nil, err
	}

	adminFee := helper.GenerateAdminFee()
	shippingFee := helper.GenerateShippingFee(address)

	// insert order to db
	order := model.Order{
		Id:              orderNumber,
		Amount:          int64(grossAmount + shippingFee + adminFee),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		PaidAt:          sql.NullTime{Valid: false},
		ExpiryDate:      paymentMethod.ExpiryTime,
		ShippingFee:     shippingFee,
		AdminFee:        adminFee,
		OrderStatus:     gopayRespondMd.TransactionStatus,
		UserId:          userId,
		AddressId:       address.Id,
		PaymentMethodId: paymentMethod.Id,
	}
	orderRespond, err := uc.insertOrder.Insert(order)
	if err != nil {
		return nil, nil, err
	}

	// add product to prodcut
	for _, productsDetail := range productsDetails {
		_, err = uc.insertProductOrderDetail.Insert(productsDetail)
		if err != nil {
			return nil, nil, err
		}
	}

	// delete all cart user
	err = uc.deleteCartByUser.Delete(userId)
	if err != nil {
		return nil, nil, err
	}

	return &orderRespond, &paymentMethod, nil
}
