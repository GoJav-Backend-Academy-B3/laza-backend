package payment_method

import "gorm.io/gorm"

type PaymentMethod struct {
	db *gorm.DB
}

func NewPaymentMethod(db *gorm.DB) *PaymentMethod {
	return &PaymentMethod{db: db}
}
