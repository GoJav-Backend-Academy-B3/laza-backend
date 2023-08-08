package creditcard

import "gorm.io/gorm"

type CreditCardRepo struct {
	db *gorm.DB
}

func NewCreditCardRepo(db *gorm.DB) *CreditCardRepo {
	return &CreditCardRepo{db: db}
}
