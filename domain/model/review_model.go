package model

import "time"

type Review struct {
	Id        uint      `json:"id,omitempty" gorm:"primarykey"`
	Comment   string    `json:"comment,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UserId    uint64    `json:"user_id,omitempty"`
	ProductId uint64    `json:"product_id,omitempty"`
}

func (Review) TableName() string {
	return "review"
}
