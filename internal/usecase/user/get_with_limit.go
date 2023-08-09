package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type GetWithLimitUserUsecase struct {
	getWithLimitAction repositories.GetWithLimitAction[model.User]
	countAction        action.Count
}

func NewGetWithLimitUserUsecase(userRepo user.UserRepo) contract.GetWithLimitUserUsecase {
	return &GetWithLimitUserUsecase{
		getWithLimitAction: &userRepo,
		countAction:        &userRepo,
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

	count, err := uc.countAction.Count()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}
	totalPage := helper.PageCount(count, int64(perpage))

	response := map[string]any{
		"meta": map[string]any{
			"page":       page,
			"perpage":    perpage,
			"total_page": totalPage,
		},
		"data": result,
	}
	return helper.GetResponse(response, 200, false)
}
