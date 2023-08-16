package payment_method

import "github.com/phincon-backend/laza/domain/model"

func (r *PaymentMethod) GetById(id any) (order model.PaymentMethod, err error) {
	tx := r.db.First(&order, "id = ?", id)
	err = tx.Error
	return
}
