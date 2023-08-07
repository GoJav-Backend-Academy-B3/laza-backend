package cart

import (
	"github.com/phincon-backend/laza/domain/response"
)

type GetCartByIdAction interface {
	GetCartById(userId any) ([]response.CartPorduct, error)
}
