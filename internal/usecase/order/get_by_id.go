package order

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	"time"
)

type GetByIdUsecase struct {
	getOrder               repositories.GetByIdAction[model.Order]
	updateOrder            repositories.UpdateAction[model.Order]
	getMidtransTransaction midtranscore.FetchMidtransTransactionAction
}

func NewGetByIdUsecase(getOrder repositories.GetByIdAction[model.Order], updateOrder repositories.UpdateAction[model.Order], getMidtransTransaction midtranscore.FetchMidtransTransactionAction) *GetByIdUsecase {
	return &GetByIdUsecase{getOrder: getOrder, updateOrder: updateOrder, getMidtransTransaction: getMidtransTransaction}
}

// Execute implements product.SearchProductByNameUsecase.
func (u *GetByIdUsecase) Execute(orderId string) (order model.Order, err error) {
	// get order
	order, err = u.getOrder.GetById(orderId)
	if err != nil {
		return
	}

	// get status transaction
	midtransTransaction, err := u.getMidtransTransaction.FetchMidtransTransaction(order.Id)

	order.OrderStatus = midtransTransaction.TransactionStatus
	order.UpdatedAt = time.Now()

	update, err := u.updateOrder.Update(orderId, order)
	if err != nil {
		return model.Order{}, err
	}

	return update, err
}
