package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *BrandRepo) GetById(id any) (brand model.Brand, err error) {
	err = r.db.Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, image_url, price")
	}).First(&brand, "id = ?", id).Error

	if err != nil {
		return
	}

	return
}