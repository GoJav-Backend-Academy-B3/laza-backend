package bank

import "gorm.io/gorm"

type BankRepo struct {
	db *gorm.DB
}

func NewBankRepo(db *gorm.DB) *BankRepo {
	return &BankRepo{db: db}
}
