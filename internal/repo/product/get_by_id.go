package product

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ProductRepo) GetById(id any) (e model.Product, err error) {
	tx := r.db.Preload("Sizes").First(&e, "id = ?", id)
	err = tx.Error
	return
}
