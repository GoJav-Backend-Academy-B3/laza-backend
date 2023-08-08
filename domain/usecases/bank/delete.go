package bank

import "github.com/phincon-backend/laza/helper"

type DeleteBankUsecase interface {
	Execute(id uint64) *helper.Response
}
