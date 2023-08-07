package product

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductRepo) Delete(id_r any) error {

	id := id_r.(uint64)
	var product model.Product

	tx := r.db.First(&product, "id = ?", id)
	// TODO: result.Error might return gorm.ErrRecordNotFound if
	//       no `id` can be found
	err := tx.Error

	tx = r.db.Delete(&product)
	err = tx.Error

	return err
}
