package order

import (
	"github.com/phincon-backend/laza/domain/model"
)

type CreateOrderWithGopayUsecase interface {
	Execute(userId uint64, addressId int, callbackUrl string) (*model.Order, *model.PaymentMethod, error)
}
