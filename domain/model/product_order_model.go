package model

type ProductOrder struct {
	ProductId uint64 `json:"product_id,omitempty" gorm:"primarykey"`
	OrderId   uint64 `json:"order_id,omitempty" gorm:"primarykey"`
	Quantity  uint16 `json:"quantity"`
	Price     float64
}
