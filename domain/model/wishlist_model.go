package model

type Wishlist struct {
	UserId    uint64 `json:"user_id,omitempty" gorm:"primarykey"`
	ProductId uint64 `json:"product_id,omitempty" gorm:"primarykey"`
	IsLiked   bool   `json:"is_liked,omitempty"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}
