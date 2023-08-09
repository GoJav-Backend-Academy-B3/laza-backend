package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories/brand"
	usecase "github.com/phincon-backend/laza/domain/usecases/brand"
	repository "github.com/phincon-backend/laza/internal/repo/brand"
)

type searchBrandByNameUsecaseImpl struct {
	findByNameAction brand.FindByNameAction
}

// Execute implements brand.SearchBrandByNameUsecase.
func (u *searchBrandByNameUsecaseImpl) Execute(brandName string) (brands []model.Brand, err error) {
	brands, err = u.findByNameAction.FindByName(brandName)
	if err != nil {
		return
	}

	return
}

func NewSearchBrandByNameUsecaseImpl(brandRepo repository.BrandRepo) usecase.SearchBrandByNameUsecase {
	return &searchBrandByNameUsecaseImpl{findByNameAction: &brandRepo}
}
