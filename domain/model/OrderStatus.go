package model

type OrderStatus struct {
	Id     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}
