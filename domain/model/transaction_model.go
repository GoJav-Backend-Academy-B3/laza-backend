package model

import (
	"database/sql"
)

type Transaction struct {
	Id                string          `json:"id,omitempty" gorm:"primarykey"`
	GrossAmount       sql.NullFloat64 `json:"gross_amount,omitempty"`
	PaymentType       sql.NullString  `json:"payment_type,omitempty"`
	Currency          sql.NullString  `json:"currency,omitempty"`
	TransactionStatus sql.NullString  `json:"transaction_status,omitempty"`
	Signature         sql.NullString  `json:"signature,omitempty"`
	FraudStatus       sql.NullString  `json:"fraud_status,omitempty"`
	ExpiryTime        sql.NullTime    `json:"expiry_time"`
	SettlementTime    sql.NullTime    `json:"settlement_time"`
	TransactionTime   sql.NullTime    `json:"transaction_time"`
	OrderId           uint64          `json:"order_id,omitempty"`
}

func (Transaction) TableName() string {
	return "order"
}
