package cart

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type InsertCartUsecase interface {
	Execute(userid uint64, rb requests.CartRequest) (_result model.Cart, err error)
}
