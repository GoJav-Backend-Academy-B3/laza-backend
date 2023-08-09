package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	midtrans_core "github.com/phincon-backend/laza/external/midtrans"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/order"
	addressRepo "github.com/phincon-backend/laza/internal/repo/address"
	gopayRepo "github.com/phincon-backend/laza/internal/repo/gopay"
	midtrans "github.com/phincon-backend/laza/internal/repo/midtrans_repo"
	orderRepo "github.com/phincon-backend/laza/internal/repo/order"
	productRepo "github.com/phincon-backend/laza/internal/repo/product"
	productOrderRepo "github.com/phincon-backend/laza/internal/repo/product_order"
	transactionBankRepo "github.com/phincon-backend/laza/internal/repo/transaction_bank"
	orderUsecase "github.com/phincon-backend/laza/internal/usecase/order"
)

func NewOrderHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs
	midtransCore := midtrans_core.Init()

	orderRepo := orderRepo.NewOrderRepo(gorm)
	addressRepo := addressRepo.NewAddressRepo(gorm)
	gopayRepo := gopayRepo.NewGopayRepo(gorm)
	productRepo := productRepo.NewProductRepo(gorm)
	productOrderRepo := productOrderRepo.NewProductOrderRepo(gorm)
	transactionBankRepo := transactionBankRepo.NewTransactionBankRepo(gorm)
	midtransRepo := midtrans.NewMidtransRepo(midtransCore)

	createOrderWithGopay := orderUsecase.NewCreateOrderWithGopayUsecase(orderRepo, addressRepo, midtransRepo, gopayRepo, orderRepo, productRepo, productOrderRepo)
	createOrderWithBank := orderUsecase.NewCreateOrderWithBankUsecase(orderRepo, addressRepo, midtransRepo, transactionBankRepo, orderRepo, productRepo, productOrderRepo)

	return handler.NewOrderHandler(createOrderWithGopay, createOrderWithBank)

}
