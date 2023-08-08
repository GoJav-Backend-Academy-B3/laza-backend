package model

type Bank struct {
	Id       uint64 `json:"id,omitempty" gorm:"primarykey" swagignore:"true"`
	BankName string `json:"bank_name,omitempty" validate:"required" form:"bank_name" `
	BankCode string `json:"bank_code,omitempty" validate:"required,numeric" form:"bank_code" `
	LogoUrl  string `json:"image,omitempty" form:"image" `
}

func (Bank) TableName() string {
	return "bank"
}
