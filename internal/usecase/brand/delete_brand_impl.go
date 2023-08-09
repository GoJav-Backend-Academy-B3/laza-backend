package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/brand"
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

func NewDeleteBrandUsecaseImpl(deleteBrandAction repositories.DeleteAction[model.Brand]) brand.DeleteBrandByIdUsecase {
	return &deleteBrandUsecaseImpl{deleteBrandAction: deleteBrandAction}
}
