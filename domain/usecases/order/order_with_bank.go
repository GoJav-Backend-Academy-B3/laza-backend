package order

import (
	"github.com/phincon-backend/laza/domain/model"
)

type CreateOrderWithBankUsecase interface {
	Execute(userId uint64, addressId int, bank string) (*model.Order, *model.PaymentMethod, error)
}
