package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/brand"
	repository "github.com/phincon-backend/laza/internal/repo/brand"
)

type deleteBrandUsecaseImpl struct {
	deleteBrandAction repositories.DeleteAction[model.Brand]
}

// Execute implements brand.DeleteBrandByIdUsecase.
func (u *deleteBrandUsecaseImpl) Execute(brandId uint64) (err error) {
	err = u.deleteBrandAction.Delete(brandId)
	if err != nil {
		return
	}

	return
}

func NewDeleteBrandUsecaseImpl(brandRepo repository.BrandRepo) brand.DeleteBrandByIdUsecase {
	return &deleteBrandUsecaseImpl{deleteBrandAction: &brandRepo}
}
