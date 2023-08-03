package model

type Bank struct {
	Id       int    `json:"id,omitempty"`
	BankName string `json:"bank-name,omitempty"`
	BankCode string `json:"bank-code,omitempty"`
	LogoUrl  string `json:"logo-url,omitempty"`
}
