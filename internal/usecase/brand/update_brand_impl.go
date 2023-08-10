package brand

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/brand"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/brand"
	"github.com/phincon-backend/laza/helper"
	repository "github.com/phincon-backend/laza/internal/repo/brand"
)

type updateBrandUsecaseImpl struct {
	updateBrandAction repositories.UpdateAction[model.Brand]
	isBrandExist      action.IsBrandExistAction
}

// Execute implements brand.UpdateBrandNameByIdUsecase.
func (u *updateBrandUsecaseImpl) Execute(id uint64, request requests.BrandRequest) (brand model.Brand, err error) {
	brandExist := u.isBrandExist.IsBrandExist(request.Name)
	if brandExist {
		return brand, errors.New("brand is exist")
	}

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

func NewUpdateBrandImpl(brandRepo repository.BrandRepo) brand.UpdateBrandNameByIdUsecase {
	return &updateBrandUsecaseImpl{
		updateBrandAction: &brandRepo,
		isBrandExist:      &brandRepo,
	}
}
