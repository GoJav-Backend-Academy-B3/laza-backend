package user

import (
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type GetAllUserUsecase struct {
	getAllAction repositories.GetAllAction[response.User]
}

func NewGetAllUserUsecase(repo repositories.GetAllAction[response.User]) user.GetAllUserUsecase {
	return &GetAllUserUsecase{getAllAction: repo}
}

// Excute implements user.GetAllUserUsecase.
func (uc *GetAllUserUsecase) Execute() *helper.Response {
	result, err := uc.getAllAction.GetAll()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
