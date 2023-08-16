package order

import (
	"database/sql"
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	midtranscore "github.com/phincon-backend/laza/domain/repositories/midtrans"
	productOrderDetailRepo "github.com/phincon-backend/laza/domain/repositories/product_order_detail"
	"time"
)

type GetByIdUsecase struct {
	getOrder                 repositories.GetByIdAction[model.Order]
	updateOrder              repositories.UpdateAction[model.Order]
	getMidtransTransaction   midtranscore.FetchMidtransTransactionAction
	getAllProductOrderDetail productOrderDetailRepo.GetAllByOrder
}

func NewGetByIdUsecase(getOrder repositories.GetByIdAction[model.Order], updateOrder repositories.UpdateAction[model.Order], getMidtransTransaction midtranscore.FetchMidtransTransactionAction, getAllProductOrderDetail productOrderDetailRepo.GetAllByOrder) *GetByIdUsecase {
	return &GetByIdUsecase{getOrder: getOrder, updateOrder: updateOrder, getMidtransTransaction: getMidtransTransaction, getAllProductOrderDetail: getAllProductOrderDetail}
}

// Execute implements product.SearchProductByNameUsecase.
func (u *GetByIdUsecase) Execute(orderId string) (order model.Order, productDetails []model.ProductOrderDetail, err error) {
	// get order
	order, err = u.getOrder.GetById(orderId)
	if err != nil {
		return
	}

	// get status transaction
	midtransTransaction, err := u.getMidtransTransaction.FetchMidtransTransaction(order.Id)

	if midtransTransaction.TransactionStatus == "success" || midtransTransaction.TransactionStatus == "settlement" || midtransTransaction.TransactionStatus == "capture" {
		// parsing time
		parsedTime, errParse := time.Parse("2006-01-02 15:04:05", midtransTransaction.TransactionTime)
		if errParse != nil {
			return order, productDetails, errParse
		}
		order.PaidAt = sql.NullTime{
			Time:  parsedTime,
			Valid: true,
		}
	}
	if order.OrderStatus != midtransTransaction.TransactionStatus {
		order.OrderStatus = midtransTransaction.TransactionStatus
		order.UpdatedAt = time.Now()
		updatedOrder, err := u.updateOrder.Update(orderId, order)
		if err != nil {
			return updatedOrder, productDetails, err
		}
		return updatedOrder, productDetails, err
	}

	productDetails, err = u.getAllProductOrderDetail.GetAllByOrder(orderId)
	if err != nil {
		return
	}

	return order, productDetails, err
}
