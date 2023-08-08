package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/brand"
	"github.com/phincon-backend/laza/helper"
)

type updateBrandUsecaseImpl struct {
	updateBrandAction repositories.UpdateAction[model.Brand]
}

// Execute implements brand.UpdateBrandNameByIdUsecase.
func (u *updateBrandUsecaseImpl) Execute(id uint64, request requests.BrandRequest) (brand model.Brand, err error) {
	file, err := request.LogoUrl.Open()
	if err != nil {
		// TODO: Should return error here
		return
	}
	defer file.Close()

	url, err := helper.UploadImageFile("brand", file)
	if err != nil {
		// TODO: Should return error here
		return
	}

	brand = model.Brand{
		Id:      id,
		Name:    request.Name,
		LogoUrl: url,
	}

	brand, err = u.updateBrandAction.Update(id, brand)
	if err != nil {
		return
	}

	return
}

func NewUpdateBrandImpl(updateBrandAction repositories.UpdateAction[model.Brand]) brand.UpdateBrandNameByIdUsecase {
	return &updateBrandUsecaseImpl{updateBrandAction: updateBrandAction}
}
