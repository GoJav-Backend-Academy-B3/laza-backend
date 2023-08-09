package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"gorm.io/gorm"
)

type GetByIdUserUsecase struct {
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewGetByIdUserUsecase(userRepo user.UserRepo) contract.GetByIdUserUsecase {
	return &GetByIdUserUsecase{
		getByIdAction: &userRepo,
	}
}

// Excute implements user.GetByIdUserUsecase.
func (uc *GetByIdUserUsecase) Execute(id uint64) *helper.Response {
	res, err := uc.getByIdAction.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.GetResponse("NotFound: data user not found", 500, true)
		}
		return helper.GetResponse(err.Error(), 500, true)
	}

	result := response.UserModelResponse(res)
	return helper.GetResponse(result, 200, false)
}
