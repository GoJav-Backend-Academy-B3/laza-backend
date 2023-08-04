package model

type OrderStatus struct {
	Id     uint64 `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}
