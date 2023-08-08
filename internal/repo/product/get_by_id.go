package product

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ProductRepo) GetProductById(id any) (e model.Product, err error) {
	tx := r.db.First(&e, "id = ?", id)
	err = tx.Error
	return
}
