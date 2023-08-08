package transaction_bank

import "gorm.io/gorm"

type TransactionBankRepo struct {
	db *gorm.DB
}

func NewTransactionBankRepo(db *gorm.DB) *TransactionBankRepo {
	return &TransactionBankRepo{db: db}
}
