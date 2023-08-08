package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/size"
)

type AddSizeUsecaseImpl struct {
	insertAction repositories.InsertAction[model.Size]
}

// Execute implements size.AddSizeUsecase.
func (u *AddSizeUsecaseImpl) Execute(name string) (m model.Size, err error) {
	m.Size = name
	m, err = u.insertAction.Insert(m)
	return
}

func NewAddSizeUsecaseImpl(insertAction repositories.InsertAction[model.Size]) size.AddSizeUsecase {
	return &AddSizeUsecaseImpl{insertAction}
}
