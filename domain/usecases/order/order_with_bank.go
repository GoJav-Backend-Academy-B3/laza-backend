package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/request"
)

type CreateOrderWithBankUsecase interface {
	Execute(userId uint64, addressId int, bank string, products []request.ProductOrder) (*model.Order, *model.TransactionBank, error)
}
