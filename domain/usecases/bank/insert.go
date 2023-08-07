package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type InsertBankUsecase interface {
	Execute(bank model.Bank) *helper.Response
}
