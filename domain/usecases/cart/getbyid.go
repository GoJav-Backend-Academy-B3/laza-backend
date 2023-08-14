package cart

import (
	"github.com/phincon-backend/laza/domain/response"
)

type GetCartByIdUsecase interface {
	Execute(userId uint64) (_result response.CartInfo, err error)
}
