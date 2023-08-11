package model

import "gorm.io/gorm"

type Brand struct {
	Id        uint64         `json:"id,omitempty" gorm:"primarykey"`
	Name      string         `json:"name,omitempty"`
	LogoUrl   string         `json:"logo_url,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" swaggertype:"primitive,integer"`
	Products  []Product      `json:"products,omitempty" gorm:"foreignkey:BrandId"`
}

func (Brand) TableName() string {
	return "brand"
}
