package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/helper"
)

type UpdateBankUsecase interface {
	Execute(id uint64, bank model.Bank) *helper.Response
}
