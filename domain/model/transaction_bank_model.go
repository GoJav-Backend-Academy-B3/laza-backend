package model

type TransactionBank struct {
	Id         uint64 `gorm:"id"`
	BankCode   string `gorm:"bank_code"`
	BillerCode string `gorm:"biller_code"`
	VANumber   string `gorm:"va_number"`
}
