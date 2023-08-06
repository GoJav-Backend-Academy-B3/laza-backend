package cart

import "gorm.io/gorm"

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{
		db: db,
	}
}
