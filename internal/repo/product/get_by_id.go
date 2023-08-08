package product

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductRepo) GetById(id string) (model.Product, error) {

	// Grab product that matches with `id`
	var product model.Product
	tx := r.db.First(&product, "id = ?", id)
	// TODO: result.Error might return gorm.ErrRecordNotFound if
	//       no `id` can be found
	err := tx.Error

	return product, err
}
