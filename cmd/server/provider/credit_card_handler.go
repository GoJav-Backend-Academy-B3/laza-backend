package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	hd "github.com/phincon-backend/laza/internal/handler/credit_card"
	rp "github.com/phincon-backend/laza/internal/repo/credit_card"
	uc "github.com/phincon-backend/laza/internal/usecase/credit_card"
)

func NewcreditCardHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()

	gorm := dbs.(*db.PsqlDB).Dbs
	validate := validator.New()

	creditCrepo := rp.NewCreditCardRepo(gorm)
	addCcUc := uc.NewaddCreditCardUsecase(creditCrepo, validate)
	updateCcUc := uc.NewupdateCreditCardUsecase(creditCrepo, validate)
	getByIdCcUc := uc.NewgetByIdCreditCardUsecase(creditCrepo)
	getAllCcUc := uc.NewgetAllCreditCardUsecase(creditCrepo)

	return hd.NewgetCreditCardHandler(
		"/credit-card",
		"/credit-card/:id",
		"/credit-card/:id",
		"/credit-card",
		addCcUc,
		updateCcUc,
		getByIdCcUc,
		getAllCcUc,
	)
}
