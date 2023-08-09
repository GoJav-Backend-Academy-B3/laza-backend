package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type CreateOrderWithBankUsecase interface {
	Execute(userId uint64, addressId int, bank string, products []requests.ProductOrder) (*model.Order, *model.TransactionBank, error)
}
