package model

type Wishlist struct {
	UserId    uint64 `json:"user_id,omitempty"`
	ProductId uint64 `json:"product_id,omitempty"`
	IsLiked   bool   `json:"is_liked,omitempty"`
}
