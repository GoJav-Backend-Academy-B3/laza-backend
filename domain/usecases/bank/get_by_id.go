package bank

import "github.com/phincon-backend/laza/helper"

type GetByIdBankUsecase interface {
	Execute(id uint64) *helper.Response
}
