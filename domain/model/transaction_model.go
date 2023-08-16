package model

//type Transaction struct {
//	Id                string          `json:"id,omitempty" gorm:"primarykey"`
//	GrossAmount       sql.NullFloat64 `json:"gross_amount,omitempty" swaggertype:"integer"`
//	PaymentType       sql.NullString  `json:"payment_type,omitempty" swaggertype:"string"`
//	Currency          sql.NullString  `json:"currency,omitempty" swaggertype:"string"`
//	TransactionStatus sql.NullString  `json:"transaction_status,omitempty" swaggertype:"string"`
//	Signature         sql.NullString  `json:"signature,omitempty" swaggertype:"string"`
//	FraudStatus       sql.NullString  `json:"fraud_status,omitempty" swaggertype:"string"`
//	ExpiryTime        sql.NullTime    `json:"expiry_time" swaggertype:"primitive,integer"`
//	SettlementTime    sql.NullTime    `json:"settlement_time" swaggertype:"primitive,integer"`
//	TransactionTime   sql.NullTime    `json:"transaction_time" swaggertype:"primitive,integer"`
//	OrderId           uint64          `json:"order_id,omitempty" swaggertype:"primitive,integer"`
//}
//
//func (Transaction) TableName() string {
//	return "transaction"
//}
