package bank

import "github.com/phincon-backend/laza/helper"

type GetAllBankUsecase interface {
	Execute() *helper.Response
}
