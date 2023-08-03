package product

import "github.com/phincon-backend/laza/domain/entities"

func (r *ProductRepo) GetAll() (es []entities.Product, err error) {
	tx := r.db.Find(&es)
	err = tx.Error
	return
}
