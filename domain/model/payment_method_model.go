package model

import "time"

type PaymentMethod struct {
	Id               uint64    `json:"id,omitempty"  gorm:"primarykey"`
	PaymentMethod    string    `json:"payment_method"`
	Deeplink         string    `json:"deeplink,omitempty"`
	QRCodeUrl        string    `json:"qr_code,omitempty"`
	Bank             string    `json:"bank,omitempty"`
	BillerCode       string    `json:"biller_code,omitempty"`
	BillKey          string    `json:"bill_key,omitempty"`
	VANumber         string    `json:"va_number,omitempty"`
	CreditCardNumber string    `json:"credit_card_number,omitempty"`
	RedirectUrl      string    `json:"redirect_url,omitempty"`
	ExpiryTime       time.Time `json:"expiry_time"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
