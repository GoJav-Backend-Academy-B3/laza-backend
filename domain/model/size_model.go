package model

type Size struct {
	Id       uint64    `json:"id,omitempty" gorm:"primarykey"`
	Size     string    `json:"size,omitempty"`
	Products []Product `json:"product" gorm:"many2many:size_product,omitempty"`
}

func (Size) TableName() string {
	return "size"
}

func (s *Size) Update(e Size) {
	s.Size = e.Size
}
