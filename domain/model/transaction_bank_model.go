package model

type TransactionBank struct {
	Id         uint64  `json:"id" gorm:"primarykey"`
	BankCode   string  `json:"bank_code"`
	BillerCode string  `json:"biller_code"`
	VANumber   string  `json:"va_number"`
	Orders     []Order `json:"orders" gorm:"foreignkey:Id"`
}

func (TransactionBank) TableName() string {
	return "transaction_bank"
}
