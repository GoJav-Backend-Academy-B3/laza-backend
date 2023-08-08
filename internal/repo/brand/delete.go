package brand

import "github.com/phincon-backend/laza/domain/model"

func (r *BrandRepo) Delete(id any) error {
	var brand model.Brand
	err := r.db.Where("id = ?", id).Delete(&brand).Error
	if err != nil {
		return err
	}

	return nil
}
