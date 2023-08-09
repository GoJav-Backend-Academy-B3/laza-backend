package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type GetAllUserUsecase struct {
	getAllAction repositories.GetAllAction[model.User]
}

func NewGetAllUserUsecase(userRepo user.UserRepo) contract.GetAllUserUsecase {
	return &GetAllUserUsecase{
		getAllAction: &userRepo,
	}
}

// Excute implements user.GetAllUserUsecase.
func (uc *GetAllUserUsecase) Execute() *helper.Response {
	res, err := uc.getAllAction.GetAll()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	var result []response.User
	for _, v := range res {
		result = append(result, *response.UserModelResponse(v))
	}
	return helper.GetResponse(result, 200, false)
}
