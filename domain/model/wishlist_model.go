package model

type Wishlist struct {
	UserId    uint64 `json:"user_id,omitempty" gorm:"primarykey" url:"userId"`
	ProductId uint64 `json:"product_id,omitempty" gorm:"primarykey" url:"productId"`
	IsLiked   bool   `json:"is_liked"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}
