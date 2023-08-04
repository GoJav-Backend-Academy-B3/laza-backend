package model

import "time"

type Gopay struct {
	Id            uint64    `json:"id,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	DeeplLink     string    `json:"deepl_link,omitempty"`
	QRCode        string    `json:"qr_code,omitempty"`
	GetStatusLink string    `json:"get_status_link,omitempty"`
	CancelLink    string    `json:"cancel_link,omitempty"`
	ExpiryDate    time.Time `json:"expiry_date"`
}
