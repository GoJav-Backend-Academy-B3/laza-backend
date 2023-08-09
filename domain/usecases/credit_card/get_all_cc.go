package credit_card

import "github.com/phincon-backend/laza/helper"

type GetAllCreditCardUsecase interface {
	Execute(userId uint64) *helper.Response
}
