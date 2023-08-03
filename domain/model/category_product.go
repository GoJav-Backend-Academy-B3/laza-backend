package model

type CategoryProduct struct {
	ProductId  uint64 `json:"product_id,omitempty"`
	CategoryId uint64 `json:"category_id,omitempty"`
}
