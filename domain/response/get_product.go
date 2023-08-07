package response

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
)

type Product struct {
	Id          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	CategoryId  uint64    `json:"category_id,omitempty"`
	BrandId     uint64    `json:"brand_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type ProductReview struct {
	Avrg_Rating float64               `json:"rating_avrg"`
	Total       int                   `json:"total"`
	Products    []model.ProductReview `json:"products"`
}

func (p *Product) FillFromEntity(e model.Product) {
	p.Id = e.Id
	p.Name = e.Name
	p.Description = e.Description
	p.ImageUrl = e.ImageUrl
	p.Price = e.Price
	p.CategoryId = e.CategoryId
	p.BrandId = e.BrandId
	p.CreatedAt = e.CreatedAt
	p.UpdatedAt = e.UpdatedAt
}
