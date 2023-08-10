package order

import (
	"errors"
	"github.com/phincon-backend/laza/domain/model"
)

func (r *OrderRepo) Update(id any, ts model.Order) (model.Order, error) {
	if id.(string) != ts.Id {
		return ts, errors.New("id not found")
	}

	tx := r.db.Save(&ts)
	err := tx.Error

	return ts, err
}
