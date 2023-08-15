package credit_card

import (
	"github.com/phincon-backend/laza/domain/model"
	cc_request "github.com/phincon-backend/laza/domain/requests"
)

type UpdateCreditCardUsecase interface {
	Execute(userId, ccId uint64, md cc_request.CreditCardRequest) (_result model.CreditCard, statusCode int, err error)
}
