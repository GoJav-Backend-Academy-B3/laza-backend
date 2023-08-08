package response

import (
	"github.com/phincon-backend/laza/domain/model"
)

type ReqReview struct {
	Products []model.ProductReview `json:"products"`
	Result   []model.ProductReview `json:"result"`
}

func (ReqReview) TableName() string {
	return "review"
}
