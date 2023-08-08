package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/size"
)

type GetAllSizeUsecaseImpl struct {
	getAllAction repositories.GetAllAction[model.Size]
}

// Execute implements size.GetAllSizeUsecase.
func (u *GetAllSizeUsecaseImpl) Execute() (models []model.Size, err error) {
	models, err = u.getAllAction.GetAll()
	return
}

func NewGetAllSizeUsecaseImpl(
	action repositories.GetAllAction[model.Size],
) size.GetAllSizeUsecase {
	return &GetAllSizeUsecaseImpl{
		getAllAction: action,
	}
}
