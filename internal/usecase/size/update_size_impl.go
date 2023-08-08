package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/size"
)

type UpdateSizeUsecaseImpl struct {
	updateAction repositories.UpdateAction[model.Size]
}

// Execute implements size.UpdateSizeUsecase.
func (u *UpdateSizeUsecaseImpl) Execute(id uint64, m model.Size) (model model.Size, err error) {

	model, err = u.updateAction.Update(id, m)
	return
}

func NewUpdateSizeUsecaseImpl(action repositories.UpdateAction[model.Size]) size.UpdateSizeUsecase {
	return &UpdateSizeUsecaseImpl{
		updateAction: action,
	}
}
