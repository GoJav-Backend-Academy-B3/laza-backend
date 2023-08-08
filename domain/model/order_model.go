package model

import (
	"database/sql"
	"time"
)

type Order struct {
	Id                string        `json:"id,omitempty"  gorm:"primarykey"`
	Amount            int64         `json:"amount"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	UserId            uint64        `json:"user_id"`
	OrderStatusId     uint64        `json:"order_status"`
	AddressId         uint64        `json:"address_id"`
	CreditCardId      sql.NullInt64 `json:"credit_card_id,omitempty"`
	TransactionBankId sql.NullInt64 `json:"transaction_bank_id,omitempty"`
	GopayId           sql.NullInt64 `json:"gopay_id,omitempty"`
	Products          []Product     `json:"products" gorm:"many2many:product_order"`
	Transactions      []Transaction `json:"transactions" gorm:"foreignkey:Id"`
}

func (Order) TableName() string {
	return "orders"
}
