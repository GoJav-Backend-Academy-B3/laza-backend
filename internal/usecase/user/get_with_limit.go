package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type GetWithLimitUserUsecase struct {
	getWithLimitAction repositories.GetWithLimitAction[model.User]
}

func NewGetWithLimitUserUsecase(userRepo user.UserRepo) contract.GetWithLimitUserUsecase {
	return &GetWithLimitUserUsecase{
		getWithLimitAction: &userRepo,
	}
}

// Excute implements user.GetWithLimitUserUsecase.
func (uc *GetWithLimitUserUsecase) Execute(page, perpage uint64) *helper.Response {
	if page == 0 || page < 1 {
		page = 1
	}

	if perpage == 0 || perpage < 1 {
		perpage = 5
	}

	offset := (page - 1) * perpage
	res, err := uc.getWithLimitAction.GetWithLimit(offset, perpage)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	var result []response.User
	for _, v := range res {
		result = append(result, *response.UserModelResponse(v))
	}
	return helper.GetResponse(result, 200, false)
}
