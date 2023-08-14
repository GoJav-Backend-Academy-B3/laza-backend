package model

type Wishlist struct {
	UserId    uint64 `json:"user_id,omitempty" gorm:"primaryKey" url:"userId"`
	ProductId uint64 `json:"product_id,omitempty" gorm:"primaryKey" url:"productId"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}
