package model

import (
	"time"
)

type Product struct {
	Id          uint64
	Name        string
	Description string
	ImageUrl    string
	Price       float32
	CategoryId  uint64
	BrandId     uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
