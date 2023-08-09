package order

import "github.com/phincon-backend/laza/domain/model"

func (r *OrderRepo) GetById(id any) (order model.Order, err error) {
	idConv := id.(string)
	tx := r.db.First(&order).Where("id = ?", idConv)
	err = tx.Error
	return
}
