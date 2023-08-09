package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type CreateOrderWithGopayUsecase interface {
	Execute(userId uint64, addressId int, callbackUrl string, products []requests.ProductOrder) (*model.Order, *model.Gopay, error)
}
