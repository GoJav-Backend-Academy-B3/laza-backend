package response

import "github.com/phincon-backend/laza/domain/model"

type GetReviews struct {
	Avrg_Rating float64               `json:"rating_avrg"`
	Total       int                   `json:"total"`
	Reviews     []model.ProductReview `json:"reviews"`
}
