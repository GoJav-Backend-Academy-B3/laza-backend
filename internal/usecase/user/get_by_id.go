package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type GetByIdUserUsecase struct {
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewGetByIdUserUsecase(repo repositories.GetByIdAction[model.User]) user.GetByIdUserUsecase {
	return &GetByIdUserUsecase{getByIdAction: repo}
}

// Excute implements user.GetByIdUserUsecase.
func (uc *GetByIdUserUsecase) Execute(id uint64) *helper.Response {
	result, err := uc.getByIdAction.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.GetResponse("NotFound: data user not found", 500, true)
		}
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
