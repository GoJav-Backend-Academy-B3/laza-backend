package sizeproduct

import "github.com/phincon-backend/laza/domain/model"

type GetSizeByProduct interface {
	GetProductById(productID uint64) ([]model.SizeProduct, error)
	GetSizeStatsByProduct(productID uint64) (uint64, string, error)
}
