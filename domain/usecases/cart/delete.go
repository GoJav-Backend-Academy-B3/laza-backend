package cart

import (
	"github.com/phincon-backend/laza/domain/requests"
)

type DeleteCartUsecase interface {
	Execute(userId uint64, rb requests.CartRequest) (value any, err error)
}
