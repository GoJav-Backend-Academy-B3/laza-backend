package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/size"
)

type DeleteSizeUsecaseImpl struct {
	deleteAction repositories.DeleteAction[model.Size]
}

// Execute implements size.DeleteSizeUsecase.
func (u *DeleteSizeUsecaseImpl) Execute(id uint64) error {
	return u.deleteAction.Delete(id)
}

func NewDeleteSizeUsecaseImpl(action repositories.DeleteAction[model.Size]) size.DeleteSizeUsecase {
	return &DeleteSizeUsecaseImpl{
		deleteAction: action,
	}
}
