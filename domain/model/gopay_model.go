package model

import "time"

type Gopay struct {
	Id            uint64    `json:"id,omitempty"  gorm:"primarykey"`
	Phone         string    `json:"phone"`
	Deeplink      string    `json:"deeplink"`
	QRCode        string    `json:"qr_code"`
	GetStatusLink string    `json:"get_status_link"`
	CancelLink    string    `json:"cancel_link"`
	ExpiryTime    time.Time `json:"expiry_time"`
	Orders        []Order   `json:"orders" gorm:"foreignkey:Id"`
}

func (Gopay) TableName() string {
	return "gopay"
}
