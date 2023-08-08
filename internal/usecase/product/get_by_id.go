package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/product"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type GetByIdProductUsecase struct {
	getProductByIdAction repositories.GetByIdAction[model.Product]
}

func NewGetByIdProductUsecase(repo repositories.GetByIdAction[model.Product]) product.GetByIdProductUsecase {
	return &GetByIdProductUsecase{getProductByIdAction: repo}
}

func (uc *GetByIdProductUsecase) Execute(id uint64) *helper.Response {
	result, err := uc.getProductByIdAction.GetById(id)
	if err != nil || err == gorm.ErrRecordNotFound {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
