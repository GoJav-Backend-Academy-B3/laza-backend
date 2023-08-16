package response

import (
	"github.com/phincon-backend/laza/domain/model"
	"time"
)

type BasicOrderResponse struct {
	Id              string    `json:"id,omitempty"`
	Amount          int64     `json:"amount,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	PaidAt          time.Time `json:"paid_at,omitempty"`
	ExpiryDate      time.Time `json:"expiry_date,omitempty"`
	ShippingFee     int       `json:"shipping_fee,omitempty"`
	AdminFee        int       `json:"admin_fee,omitempty"`
	OrderStatus     string    `json:"order_status,omitempty"`
	UserId          uint64    `json:"user_id,omitempty"`
	AddressId       uint64    `json:"address_id,omitempty"`
	PaymentMethodId uint64    `json:"payment_method_id,omitempty"`
}

func (r *BasicOrderResponse) FillFromEntity(m *model.Order) {
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
}
