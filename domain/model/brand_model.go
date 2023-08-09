package model

type Brand struct {
	Id       uint64    `json:"id,omitempty" gorm:"primarykey"`
	Name     string    `json:"name,omitempty"`
	LogoUrl  string    `json:"logo_url,omitempty"`
	Products []Product `json:"products,omitempty"gorm:"foreignkey:BrandId;references:Id`
}

func (Brand) TableName() string {
	return "brand"
}
