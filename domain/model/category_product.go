package model

type CategoryProduct struct {
	ProductId  uint64 `json:"product_id,omitempty" gorm:"primarykey"`
	CategoryId uint64 `json:"category_id,omitempty" gorm:"primarykey"`
}

func (CategoryProduct) TableName() string {
	return "category_product"
}
