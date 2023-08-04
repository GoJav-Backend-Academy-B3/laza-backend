package model

type Category struct {
	Id       uint64    `json:"id,omitempty" gorm:"primarykey"`
	Category string    `json:"category,omitempty"`
	Products []Product `json:"products" gorm:"many2many:category_product"`
}

func (Category) TableName() string {
	return "category"
}
