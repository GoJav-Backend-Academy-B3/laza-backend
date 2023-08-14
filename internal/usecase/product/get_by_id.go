package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	review "github.com/phincon-backend/laza/domain/repositories/Review"
	action "github.com/phincon-backend/laza/domain/repositories/category"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/product"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type GetByIdProductUsecase struct {
	getProductByIdAction  repositories.GetByIdAction[model.Product]
	getByIdAction         repositories.GetByIdAction[model.Size]
	getReviewByProduct    review.GetReviewByProduct
	getCategoryByIdAction action.FindByIdAction
}

func NewGetByIdProductUsecase(product repositories.GetByIdAction[model.Product], repoReview review.GetReviewByProduct, size repositories.GetByIdAction[model.Size], category action.FindByIdAction) product.GetByIdProductUsecase {
	return &GetByIdProductUsecase{getProductByIdAction: product, getReviewByProduct: repoReview, getByIdAction: size, getCategoryByIdAction: category}
}

func (uc *GetByIdProductUsecase) Execute(id uint64) *helper.Response {
	result, err := uc.getProductByIdAction.GetById(id)

	pd := response.ProductDetail{}.FillFromEntity(result)
	review, err := uc.getReviewByProduct.GetProductById(id)
	records, err := uc.getCategoryByIdAction.FindById(result.CategoryId)
	if err != nil || err == gorm.ErrRecordNotFound {
		return helper.GetResponse(err.Error(), 400, true)
	}
	pd.Categories = records
	if len(review) >= 1 {
		pd.Reviews = review[:1]
	} else {
		pd.Reviews = review
	}

	return helper.GetResponse(pd, 200, false)
}
