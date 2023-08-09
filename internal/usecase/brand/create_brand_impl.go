package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/brand"
	"github.com/phincon-backend/laza/helper"
)

type createBrandUsecaseImpl struct {
	insertCategoryAction repositories.InsertAction[model.Brand]
}

// Execute implements brand.CreateBrandUsecase.
func (u *createBrandUsecaseImpl) Execute(request requests.BrandRequest) (brand model.Brand, err error) {

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
		Name:    request.Name,
		LogoUrl: url,
	}

	brand, err = u.insertCategoryAction.Insert(brand)
	if err != nil {
		return
	}

	return
}

func NewCreateBrandUseCaseImpl(insertCategoryAction repositories.InsertAction[model.Brand]) brand.CreateBrandUsecase {
	return &createBrandUsecaseImpl{insertCategoryAction: insertCategoryAction}
}
