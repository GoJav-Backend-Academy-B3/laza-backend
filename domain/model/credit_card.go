package model

type CreditCard struct {
	Id           uint    `json:"id,omitempty"  gorm:"primarykey"`
	CardNumber   string  `json:"card_number,omitempty"`
	ExpiredMonth string  `json:"expired_month,omitempty"`
	ExpiredYear  string  `json:"expired_year,omitempty"`
	UserId       uint64  `json:"user_id,omitempty"`
	Orders       []Order `json:"orders" gorm:"foreignkey:Id"`
}

func (CreditCard) TableName() string {
	return "credit_card"
}
