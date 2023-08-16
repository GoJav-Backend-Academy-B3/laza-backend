package response

import (
	"github.com/phincon-backend/laza/domain/model"
	"time"
)

type Order struct {
	Id              string    `json:"id,omitempty"`
	Amount          int64     `json:"amount"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PaidAt          time.Time `json:"paid_at,omitempty"`
	ExpiryDate      time.Time `json:"expiry_date,omitempty"`
	ShippingFee     int       `json:"shipping_fee"`
	AdminFee        int       `json:"admin_fee"`
	OrderStatus     string    `json:"order_status"`
	UserId          uint64    `json:"user_id,omitempty"`
	AddressId       uint64    `json:"address_id"`
	PaymentMethodId uint64    `json:"payment_method_id"`
}

func (r *Order) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.OrderStatus = m.OrderStatus
	r.AddressId = m.AddressId
	r.PaidAt = m.PaidAt.Time
	r.AdminFee = m.AdminFee
	r.ExpiryDate = m.ExpiryDate
	r.PaymentMethodId = m.PaymentMethodId
	r.ShippingFee = m.ShippingFee
	//r.GopayId = m.GopayId.Int64
	//r.CreditCardId = m.CreditCardId.Int64
	//r.TransactionBankId = m.TransactionBankId.Int64
}
