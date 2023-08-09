package order

import "github.com/phincon-backend/laza/domain/model"

func (r *OrderRepo) GetById(id any) (order model.Order, err error) {
	tx := r.db.First(&order, "id = ?", id)
	err = tx.Error
	return
}
