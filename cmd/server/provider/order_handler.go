package provider

import (
	midtrans_core "github.com/phincon-backend/laza/config"
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/order"
	addressRepo "github.com/phincon-backend/laza/internal/repo/address"
	brandRepo "github.com/phincon-backend/laza/internal/repo/brand"
	cartRepo "github.com/phincon-backend/laza/internal/repo/cart"
	categoryRepo "github.com/phincon-backend/laza/internal/repo/category"
	creditCardRepo "github.com/phincon-backend/laza/internal/repo/credit_card"
	midtrans "github.com/phincon-backend/laza/internal/repo/midtrans_repo"
	orderRepo "github.com/phincon-backend/laza/internal/repo/order"
	paymentMethodRepo "github.com/phincon-backend/laza/internal/repo/payment_method"
	productRepo "github.com/phincon-backend/laza/internal/repo/product"
	productOrderDetailRepo "github.com/phincon-backend/laza/internal/repo/product_order_detail"
	orderUsecase "github.com/phincon-backend/laza/internal/usecase/order"
)

func NewOrderHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs
	midtransCore := midtrans_core.Init()

	orderRepo := orderRepo.NewOrderRepo(gorm)
	addressRepo := addressRepo.NewAddressRepo(gorm)
	productRepo := productRepo.NewProductRepo(gorm)

	creditCardRepo := creditCardRepo.NewCreditCardRepo(gorm)
	productOrderDetailRepo := productOrderDetailRepo.NewProductOrderDetailRepo(gorm)
	cartRepo := cartRepo.NewCartRepo(gorm)
	brandRepo := brandRepo.NewBrandRepo(gorm)
	categoryRepo := categoryRepo.NewCategoryRepo(gorm)
	paymentMethodRepo := paymentMethodRepo.NewPaymentMethod(gorm)

	midtransRepo := midtrans.NewMidtransRepo(midtransCore)

	createOrderWithGopay := orderUsecase.NewCreateOrderWithGopayUsecase(
		orderRepo,
		addressRepo,
		midtransRepo,
		orderRepo,
		productRepo,
		productOrderDetailRepo,
		categoryRepo,
		brandRepo,
		paymentMethodRepo,
		cartRepo,
		cartRepo,
	)

	createOrderWithBank := orderUsecase.NewCreateOrderWithBankUsecase(orderRepo,
		addressRepo,
		midtransRepo,
		orderRepo,
		productRepo,
		productOrderDetailRepo,
		categoryRepo,
		brandRepo,
		paymentMethodRepo,
		cartRepo,
		cartRepo,
	)

	createOrderWithCC := orderUsecase.NewCreateOrderWithCCUsecase(
		orderRepo,
		addressRepo,
		midtransRepo,
		midtransRepo,
		creditCardRepo,
		orderRepo,
		productRepo,
		productOrderDetailRepo,
		categoryRepo,
		brandRepo,
		paymentMethodRepo,
		cartRepo,
		cartRepo,
	)

	getOrderById := orderUsecase.NewGetByIdUsecase(orderRepo, orderRepo, midtransRepo, productOrderDetailRepo)
	getAllOrderByUser := orderUsecase.NewGetAllOrderByUserUsecaseImpl(orderRepo)

	return handler.NewOrderHandler(createOrderWithGopay, createOrderWithBank, createOrderWithCC, getOrderById, getAllOrderByUser)
}
