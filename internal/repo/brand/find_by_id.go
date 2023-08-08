package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *BrandRepo) GetById(id uint64) (brand model.Brand, err error) {
	err = r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, image_url, price")
	}).First(&brand, "id = ?", id).Error

	if err != nil {
		return
	}

	return
}
