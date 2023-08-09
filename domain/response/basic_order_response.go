package response

import (
	"time"
)

type BasicOrderResponse struct {
	Id          string    `json:"id,omitempty"`
	Amount      int64     `json:"amount,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserId      uint64    `json:"user_id,omitempty"`
	OrderStatus string    `json:"order_status,omitempty"`
	AddressId   uint64    `json:"address_id,omitempty"`
}
