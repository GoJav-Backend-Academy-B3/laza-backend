package order

import "github.com/phincon-backend/laza/domain/model"

func (r *OrderRepo) GetAllByUser(userId uint64) (order []model.Order, err error) {
	tx := r.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&order)
	err = tx.Error
	return
}
