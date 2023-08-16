package order

import (
	"github.com/phincon-backend/laza/domain/model"
)

type CreateOrderWithCCUsecase interface {
	Execute(userId uint64, addressId int, cc model.CreditCard, cvv string) (*model.Order, *model.PaymentMethod, error)
}
