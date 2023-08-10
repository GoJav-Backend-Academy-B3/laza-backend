package brand

import "github.com/phincon-backend/laza/domain/model"

func (r *BrandRepo) IsBrandExist(name string) bool {
	var brand model.Brand
	err := r.db.First(&brand, "name = ?", name).Error

	return err == nil
}
