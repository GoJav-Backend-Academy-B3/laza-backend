package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type GetByIdUserUsecase struct {
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewGetByIdUserUsecase(repo repositories.GetByIdAction[model.User]) user.GetByIdUserUsecase {
	return &GetByIdUserUsecase{getByIdAction: repo}
}

// Excute implements user.GetByIdUserUsecase.
func (uc *GetByIdUserUsecase) Excute(id uint64) *helper.Response {
	result, err := uc.getByIdAction.GetById(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}