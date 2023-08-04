package model

import (
	"time"
)

type User struct {
	Id                 uint                `json:"id,omitempty" gorm:"primarykey"`
	Username           string              `json:"username,omitempty"`
	Password           string              `json:"password,omitempty"`
	Email              string              `json:"email,omitempty"`
	FullName           string              `json:"full_name,omitempty"`
	IsVerified         bool                `json:"is_verified,omitempty"`
	IsAdmin            bool                `json:"is_admin,omitempty"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	Addresses          []Address           `json:"addresses" gorm:"foreignkey:Id"`
	CreditCards        []CreditCard        `json:"credit_cards" gorm:"foreignKey:Id"`
	Reviews            []Review            `json:"reviews" gorm:"foreignkey:Id"`
	CartProducts       []Product           `json:"cart_products" gorm:"many2many:cart"`
	WishlistedProducts []Product           `json:"wishlisted_product" gorm:"many2many:wishlist"`
	Orders             []Order             `json:"orders" gorm:"foreignkey:Id"`
	VerificationCodes  []VerificationCode  `json:"verification_codes" gorm:"foreignkey:Id"`
	VerificationTokens []VerificationToken `json:"verification_tokens" gorm:"foreignkey:Id"`
}
