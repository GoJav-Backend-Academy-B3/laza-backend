package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type CreateOrderWithCCUsecase interface {
	Execute(userId uint64, addressId int, cc model.CreditCard, cvv string, products []requests.ProductOrder) (*model.Order, *model.CreditCard, error)
}
