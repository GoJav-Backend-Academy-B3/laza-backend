package model

import (
	"database/sql"
	"time"
)

type Order struct {
	Id                string        `json:"id,omitempty"  gorm:"primarykey"`
	Amount            int64         `json:"amount,omitempty"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	UserId            sql.NullInt64 `json:"user_id,omitempty"`
	OrderStatusId     sql.NullInt64 `json:"order_status,omitempty"`
	AddressId         sql.NullInt64 `json:"address_id,omitempty"`
	CreditCardId      sql.NullInt64 `json:"credit_card_id,omitempty"`
	TransactionBankId sql.NullInt64 `json:"transaction_bank_id,omitempty"`
	GopayId           sql.NullInt64 `json:"gopay_id,omitempty"`
	Products          []Product     `json:"products" gorm:"many2many:product_order"`
	Transactions      []Transaction `json:"transactions" gorm:"foreignkey:Id"`
}

func (Order) TableName() string {
	return "orders"
}
