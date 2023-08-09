package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	review "github.com/phincon-backend/laza/domain/repositories/Review"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/product"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type GetByIdProductUsecase struct {
	getProductByIdAction repositories.GetByIdAction[model.Product]
	getReviewByProduct   review.GetReviewByProduct
}

func NewGetByIdProductUsecase(repo repositories.GetByIdAction[model.Product], repoReview review.GetReviewByProduct) product.GetByIdProductUsecase {
	return &GetByIdProductUsecase{getProductByIdAction: repo, getReviewByProduct: repoReview}
}

func (uc *GetByIdProductUsecase) Execute(id uint64) *helper.Response {
	result, err := uc.getProductByIdAction.GetById(id)
	if err != nil || err == gorm.ErrRecordNotFound {
		return helper.GetResponse(err.Error(), 400, true)
	}
	pd := response.ProductDetail{}.FillFromEntity(result)
	review, err := uc.getReviewByProduct.GetProductById(id)
	pd.Reviews = review

	return helper.GetResponse(pd, 200, true)
}
