package cart

import (
	"github.com/phincon-backend/laza/domain/requests"
)

type UpdateCartUsecase interface {
	Execute(userId uint64, rb requests.CartRequest) (_result any, err error)
}
