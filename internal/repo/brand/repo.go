package brand

import "gorm.io/gorm"

type BrandRepo struct {
	db *gorm.DB
}

func NewBrandRepo(db *gorm.DB) *BrandRepo {
	return &BrandRepo{
		db: db,
	}
}
