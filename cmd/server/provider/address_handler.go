package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/address"
	repository "github.com/phincon-backend/laza/internal/repo/address"
	usecase "github.com/phincon-backend/laza/internal/usecase/address"
)

func NewAddressesHandler() handlers.HandlerInterface {
	database := db.GetPostgreSQLConnection()
	gorm := database.(*db.PsqlDB).Dbs

	validate := validator.New()

	addressRepo := repository.NewAddressRepo(gorm)

	deleteUsecase := usecase.NewDeleteAddressUsecase(addressRepo, addressRepo)
	updateUsecase := usecase.NewUpdateAddressUsecase(addressRepo, addressRepo)
	getUsecase := usecase.NewGetAddrressUsecase(addressRepo, addressRepo)
	addUsecase := usecase.NewAddAddressUsecase(addressRepo, addressRepo)

	return handler.NewAddressHandler(addUsecase, getUsecase, updateUsecase, deleteUsecase, validate)
}
