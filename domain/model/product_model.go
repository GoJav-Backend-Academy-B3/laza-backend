package model

import "time"

type Product struct {
	Id          uint64    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	ImageUrl    string    `json:"image_url,omitempty"`
	Price       float32   `json:"price,omitempty"`
	CategoryId  uint64    `json:"category_id,omitempty"`
	BrandId     uint64    `json:"brand_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
