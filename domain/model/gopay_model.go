package model

import "time"

type Gopay struct {
	Id            int       `json:"id,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	DeeplLink     string    `json:"deepl-link,omitempty"`
	QRCode        string    `json:"qr-code,omitempty"`
	GetStatusLink string    `json:"get-status-link,omitempty"`
	CancelLink    string    `json:"cancel-link,omitempty"`
	ExpiryDate    time.Time `json:"expiry-date"`
}
