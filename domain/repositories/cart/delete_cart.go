package cart

import "github.com/phincon-backend/laza/domain/model"

type DeleteCartAction interface {
	DeleteCart(model model.Cart) (value any, err error)
}
