package product_order

import "gorm.io/gorm"

type ProductOrderRepo struct {
	db *gorm.DB
}

func NewProductOrderRepo(db *gorm.DB) *ProductOrderRepo {
	return &ProductOrderRepo{db: db}
}
