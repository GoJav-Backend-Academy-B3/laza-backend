package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *BrandRepo) FindAll() (brand model.Brand, err error) {
	err = r.db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, description, image_url, price")
	}).Find(&brand).Error
	if err != nil {
		return
	}

	return
}
