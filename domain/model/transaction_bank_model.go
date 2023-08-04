package model

type TransactionBank struct {
	Id         uint64  `gorm:"id" gorm:"primarykey"`
	BankCode   string  `gorm:"bank_code"`
	BillerCode string  `gorm:"biller_code"`
	VANumber   string  `gorm:"va_number"`
	Orders     []Order `json:"orders" gorm:"foreignkey:Id"`
}
