package order

import "github.com/phincon-backend/laza/domain/model"

func (r *OrderRepo) Insert(order model.Order) (model.Order, error) {
	tx := r.db.Create(&order)
	err := tx.Error

	return order, err
}
