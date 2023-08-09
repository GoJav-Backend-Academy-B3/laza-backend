package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/brand"
)

type viewBrandUsecaseImpl struct {
	findAllAction repositories.GetAllAction[model.Brand]
}

// Execute implements brand.ViewBrandUsecase.
func (u *viewBrandUsecaseImpl) Execute() (brands []model.Brand, err error) {
	brands, err = u.findAllAction.GetAll()
	if err != nil {
		return
	}

	return
}

func NewViewBrandUsecaseImpl(findAllAction repositories.GetAllAction[model.Brand]) brand.ViewBrandUsecase {
	return &viewBrandUsecaseImpl{findAllAction: findAllAction}
}
