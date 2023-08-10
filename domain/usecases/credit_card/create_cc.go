package credit_card

import (
	cc_request "github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/helper"
)

type AddCreditCardUsecase interface {
	Execute(userId uint64, md cc_request.CreditCardRequest) *helper.Response
}
