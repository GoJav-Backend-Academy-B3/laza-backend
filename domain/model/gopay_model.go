package model

import "time"

type Gopay struct {
	Id            uint64    `json:"id,omitempty"  gorm:"primarykey"`
	Phone         string    `json:"phone,omitempty"`
	DeepLink      string    `json:"deep_link,omitempty"`
	QRCode        string    `json:"qr_code,omitempty"`
	GetStatusLink string    `json:"get_status_link,omitempty"`
	CancelLink    string    `json:"cancel_link,omitempty"`
	ExpiryDate    time.Time `json:"expiry_date"`
	Orders        []Order   `json:"orders" gorm:"foreignkey:Id"`
}
