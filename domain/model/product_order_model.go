package model

type ProductOrder struct {
	ProductId uint64 `json:"product_id,omitempty" gorm:"primarykey"`
	OrderId   string `json:"order_id,omitempty" gorm:"primarykey"`
	Quantity  uint16 `json:"quantity"`
	Price     float64
}

func (ProductOrder) TableName() string {
	return "product_order"
}
