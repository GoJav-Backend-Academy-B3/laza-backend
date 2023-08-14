package payment_method

import "github.com/phincon-backend/laza/domain/model"

func (r *PaymentMethod) Insert(model model.PaymentMethod) (model.PaymentMethod, error) {
	tx := r.db.Create(&model)
	err := tx.Error

	return model, err
}
