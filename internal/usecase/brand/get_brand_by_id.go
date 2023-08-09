package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/brand"
)

type getBrandByIdUsecaseImpl struct {
	getBrandByIdAction repositories.GetByIdAction[model.Brand]
}

// Execute implements brand.GetBrandByIdUsecase.
func (u *getBrandByIdUsecaseImpl) Execute(brandId uint64) (brand model.Brand, err error) {
	brand, err = u.getBrandByIdAction.GetById(brandId)
	if err != nil {
		return
	}

	return
}

func NewGetBrandByIdUsecaseImpl(getBrandByIdAction repositories.GetByIdAction[model.Brand]) brand.GetBrandByIdUsecase {
	return &getBrandByIdUsecaseImpl{getBrandByIdAction: getBrandByIdAction}
}
