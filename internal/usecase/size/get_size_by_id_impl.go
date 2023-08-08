package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/size"
)

type GetSizeByIdUsecaseImpl struct {
	getByIdAction repositories.GetByIdAction[model.Size]
}

// Execute implements size.GetSizeById.
func (u *GetSizeByIdUsecaseImpl) Execute(id uint64) (m model.Size, err error) {
	m, err = u.getByIdAction.GetById(id)
	return
}

func NewGetSizeByIdUsecaseImpl(action repositories.GetByIdAction[model.Size]) size.GetSizeById {
	return &GetSizeByIdUsecaseImpl{
		getByIdAction: action,
	}
}
