package model

import "time"

type Cart struct {
	UserId    uint64    `json:"user_id,omitempty"`
	ProductId uint64    `json:"product_id,omitempty"`
	Quantity  int       `json:"quantity,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
