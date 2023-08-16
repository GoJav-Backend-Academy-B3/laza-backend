package product_order_detail

import (
	"gorm.io/gorm"
)

type ProductOrderDetailRepo struct {
	db *gorm.DB
}

func NewProductOrderDetailRepo(db *gorm.DB) *ProductOrderDetailRepo {
	return &ProductOrderDetailRepo{db: db}
}
