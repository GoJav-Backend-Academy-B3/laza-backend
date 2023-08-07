package model

import (
	"time"
)

type Order struct {
	Id                string        `json:"id,omitempty"  gorm:"primarykey"`
	Amount            float64       `json:"amount,omitempty"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	UserId            uint64        `json:"user_id,omitempty"`
	OrderStatus       uint64        `json:"order_status,omitempty"`
	AddressId         uint64        `json:"address_id,omitempty"`
	CreditCardId      uint64        `json:"credit_card_id,omitempty"`
	TransactionBankId uint64        `json:"transaction_bank_id,omitempty"`
	GopayId           uint64        `json:"gopay_id,omitempty"`
	Products          []Product     `json:"products" gorm:"many2many:product_order"`
	Transactions      []Transaction `json:"transactions" gorm:"foreignkey:Id"`
}

func (Order) TableName() string {
	return "orders"
}
