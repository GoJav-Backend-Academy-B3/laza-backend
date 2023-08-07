package product

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ProductRepo) Update(id_r any, e model.Product) (model.Product, error) {

	id := id_r.(uint64)
	// Grab product that matches with `id`
	var product model.Product
	tx := r.db.First(&product, "id = ?", id)
	// TODO: result.Error might return gorm.ErrRecordNotFound if
	//       no `id` can be found
	err := tx.Error

	// Modify data
	product.Update(e)

	// Update data. Please keep in mind that this Save function
	// inserts data if no matching primary key found
	tx = r.db.Save(&product)
	err = tx.Error

	return product, err
}
