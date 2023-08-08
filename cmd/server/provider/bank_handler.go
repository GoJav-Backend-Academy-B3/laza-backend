package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/bank"
	repoBank "github.com/phincon-backend/laza/internal/repo/bank"
	usecase "github.com/phincon-backend/laza/internal/usecase/bank"
)

func NewBankHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	repoBank := repoBank.NewBankRepo(gorm)
	validate := validator.New()

	getAllBank := usecase.NewGetAllBankUsecase(repoBank)
	getByIdBank := usecase.NewGetByIdBankUsecase(repoBank)
	updateBank := usecase.NewUpdateBankUsecase(repoBank)
	deleteBank := usecase.NewDeleteBankUsecase(repoBank)
	insertBank := usecase.NewInsertBankUsecase(repoBank, repoBank)

	// insertBanks := usecase.NewInsertBankUsecase(repoBanks)

	return handler.NewBankHandler(getAllBank, getByIdBank, insertBank, updateBank, deleteBank, validate)
}
