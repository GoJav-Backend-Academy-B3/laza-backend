package model

type Size struct {
	Id       uint64    `json:"id,omitempty" gorm:"primarykey"`
	Size     string    `json:"size,omitempty"`
	Products []Product `json:"product" gorm:"many2many:size_product"`
}

func (Size) TableName() string {
	return "size"
}
