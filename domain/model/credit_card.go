package model

type CreditCard struct {
	Id           uint   `json:"id,omitempty"`
	CardNumber   string `json:"card_number,omitempty"`
	ExpiredMonth string `json:"expired_month,omitempty"`
	ExpiredYear  string `json:"expired_year,omitempty"`
	UserId       uint64 `json:"user_id,omitempty"`
}
