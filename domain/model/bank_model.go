package model

type Bank struct {
	Id       uint64 `json:"id,omitempty" gorm:"primarykey"`
	BankName string `json:"bank_name,omitempty"`
	BankCode string `json:"bank_code,omitempty"`
	LogoUrl  string `json:"logo_url,omitempty"`
}

func (Bank) TableName() string {
	return "bank"
}
