package model

import "time"

type Review struct {
	Id        uint      `json:"id,omitempty" gorm:"primarykey"`
	Comment   string    `json:"comment,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	// ProducReviews []ProductReview `json:"product_reviews"`
	UserId    uint64 `json:"user_id,omitempty"`
	ProductId uint64 `json:"product_id,omitempty"`
}
type ProductReview struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Comment   string    `json:"comment"`
	Rating    float32   `json:"rating"`
	FullName  string    `json:"full_name"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (Review) TableName() string {
	return "review"
}
