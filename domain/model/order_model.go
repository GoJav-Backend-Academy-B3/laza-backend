package model

import "time"

type Order struct {
	Id                uint64    `json:"id,omitempty"`
	Amount            int       `json:"amount,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	UserId            uint64    `json:"user_id,omitempty"`
	OrderStatus       string    `json:"order_status,omitempty"`
	AddressId         uint64    `json:"address_id,omitempty"`
	CreditCardId      uint64    `json:"credit_card_id,omitempty"`
	TransactionBankId uint64    `json:"transaction_bank_id,omitempty"`
	GopayId           uint64    `json:"gopay_id,omitempty"`
}
