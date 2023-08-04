package model

import "time"

type Product struct {
	Id            uint64     `json:"id,omitempty" gorm:"primarykey"`
	Name          string     `json:"name,omitempty"`
	Description   string     `json:"description,omitempty"`
	ImageUrl      string     `json:"image_url,omitempty"`
	Price         float64    `json:"price,omitempty"`
	CategoryId    uint64     `json:"category_id,omitempty"`
	BrandId       uint64     `json:"brand_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Sizes         []Size     `json:"size" gorm:"many2many:size_product"`
	Reviews       []Review   `json:"reviews" gorm:"foreignkey:Id"`
	Categorys     []Category `json:"categorys" gorm:"many2many:category_product"` //many 2 many
	CartUsers     []User     `json:"cart_users" gorm:"many2many:cart"`
	WishlistUsers []User     `json:"wishlisted_users" gorm:"many2many:wishlist"`
	Orders        []Order    `json:"orders" gorm:"many2many:product_order"`
}

func (Product) TableName() string {
	return "product"
}