package bank

import (
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type UpdateBankUsecase interface {
	Execute(id uint64, request requests.BankRequest) *helper.Response
}
