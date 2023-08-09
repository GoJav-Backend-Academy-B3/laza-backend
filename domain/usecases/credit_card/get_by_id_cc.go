package credit_card

import "github.com/phincon-backend/laza/helper"

type GetByIdCreditCardUsecase interface {
	Execute(ccId uint64) *helper.Response
}
