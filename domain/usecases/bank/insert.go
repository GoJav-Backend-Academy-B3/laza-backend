package bank

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type InsertBankUsecase interface {
	Execute(bank requests.BankRequest) *helper.Response
}
