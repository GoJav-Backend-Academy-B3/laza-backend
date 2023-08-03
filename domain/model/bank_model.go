package model

type Bank struct {
	Id       uint64 `json:"id,omitempty"`
	BankName string `json:"bank_name,omitempty"`
	BankCode string `json:"bank_code,omitempty"`
	LogoUrl  string `json:"logo_url,omitempty"`
}
