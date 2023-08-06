package user

import (
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type GetWithLimitUserUsecase struct {
	getWithLimitAction repositories.GetWithLimitAction[response.User]
}

func NewGetWithLimitUserUsecase(repo repositories.GetWithLimitAction[response.User]) user.GetWithLimitUserUsecase {
	return &GetWithLimitUserUsecase{getWithLimitAction: repo}
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
	result, err := uc.getWithLimitAction.GetWithLimit(offset, perpage)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
