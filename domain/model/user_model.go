package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                 uint                `json:"id,omitempty" gorm:"primarykey"`
	Username           string              `json:"username,omitempty"`
	Password           string              `json:"password,omitempty"`
	Email              string              `json:"email,omitempty"`
	FullName           string              `json:"full_name,omitempty"`
	ImageUrl           string              `json:"image_url,omitempty"`
	IsVerified         bool                `json:"is_verified,omitempty"`
	IsAdmin            bool                `json:"is_admin,omitempty"`
	CreatedAt          time.Time           `json:"created_at,omitempty"`
	UpdatedAt          time.Time           `json:"updated_at,omitempty"`
	DeletedAt          gorm.DeletedAt      `json:"deleted_at,omitempty" swaggertype:"primitive,integer"`
	Addresses          []Address           `json:"addresses,omitempty" gorm:"foreignKey:UserId"`
	CreditCards        []CreditCard        `json:"credit_cards,omitempty" gorm:"foreignKey:Id"`
	Reviews            []Review            `json:"reviews,omitempty" gorm:"foreignkey:Id"`
	CartProducts       []Product           `json:"cart_products,omitempty" gorm:"many2many:cart"`
	WishlistedProducts []Product           `json:"wishlisted_product,omitempty" gorm:"many2many:wishlist"`
	Orders             []Order             `json:"orders,omitempty" gorm:"foreignkey:Id"`
	VerificationCodes  []VerificationCode  `json:"verification_codes,omitempty" gorm:"foreignkey:Id"`
	VerificationTokens []VerificationToken `json:"verification_tokens,omitempty" gorm:"foreignkey:Id"`
}

func (User) TableName() string {
	return "users"
}
