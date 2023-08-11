package model

type Category struct {
	Id       uint64    `json:"id,omitempty" gorm:"primarykey"`
	Category string    `json:"category,omitempty"`
	Products []Product `json:"products,omitempty"`
}

func (*Category) TableName() string {
	return "category"
}

// Update only update category's name.
func (c *Category) Update(other Category) {
	c.Category = other.Category
}

func (c *Category) SetCategory(category string) {
	c.Category = category
}

func (c *Category) GetCategory() string {
	return c.Category
}

func (c *Category) SetId(id uint64) {
	c.Id = id
}
