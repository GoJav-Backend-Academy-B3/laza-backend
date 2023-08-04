package model

type SizeProduct struct {
	ProductId uint64 `json:"product_id,omitempty" gorm:"primarykey"`
	SizeId    uint64 `json:"size_id,omitempty" gorm:"primarykey"`
}
