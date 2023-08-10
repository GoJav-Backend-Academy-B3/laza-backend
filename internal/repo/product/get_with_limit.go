package product

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ProductRepo) GetWithLimit(offset, limit uint64) (es []model.Product, err error) {
	tx := r.db.
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&es)
	err = tx.Error
	return
}
