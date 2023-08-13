package cart

import "github.com/phincon-backend/laza/domain/model"

type UpdateCartAction interface {
	UpdateCart(model model.Cart) (rs any, err error)
}
