package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type GetWithLimitUserUsecase struct {
	getWithLimitAction repositories.GetWithLimitAction[model.User]
}

func NewGetWithLimitUserUsecase(repo repositories.GetWithLimitAction[model.User]) user.GetWithLimitUserUsecase {
	return &GetWithLimitUserUsecase{getWithLimitAction: repo}
}

// Excute implements user.GetWithLimitUserUsecase.
func (uc *GetWithLimitUserUsecase) Excute(offset, limit uint64) *helper.Response {
	result, err := uc.getWithLimitAction.GetWithLimit(offset, limit)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
