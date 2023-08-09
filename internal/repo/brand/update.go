package brand

import "github.com/phincon-backend/laza/domain/model"

func (r *BrandRepo) Update(id any, brand model.Brand) (model.Brand, error) {
	err := r.db.Model(&brand).Where("id = ?", id).Updates(&brand).Error
	if err != nil {
		return brand, err
	}

	return brand, nil
}
