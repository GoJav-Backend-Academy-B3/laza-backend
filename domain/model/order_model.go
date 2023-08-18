package model

import (
	"database/sql"
	"time"
)

type Order struct {
	Id              string       `json:"id,omitempty"  gorm:"primarykey"`
	Amount          int64        `json:"amount"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	PaidAt          sql.NullTime `json:"paid_at" swaggertype:"primitive,string"`
	ExpiryDate      time.Time    `json:"expiry_date"`
	ShippingFee     int          `json:"shipping_fee"`
	AdminFee        int          `json:"admin_fee"`
	OrderStatus     string       `json:"order_status"`
	UserId          uint64       `json:"user_id"`
	AddressId       uint64       `json:"address_id"`
	PaymentMethodId uint64       `json:"payment_method_id"`
}

func (Order) TableName() string {
	return "orders"
}
