package order

import "github.com/phincon-backend/laza/domain/model"

func (r *OrderRepo) GetAll() (order []model.Order, err error) {
	tx := r.db.Find(&order)
	err = tx.Error
	return
}
