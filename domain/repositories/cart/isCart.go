package cart

import "github.com/phincon-backend/laza/domain/model"

type IsCarttByIdAction interface {
	IsCart(cart model.Cart) (rs bool)
}
