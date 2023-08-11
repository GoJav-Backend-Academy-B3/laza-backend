package auth

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"gorm.io/gorm"
)

type RefreshTokenUsecase struct {
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewRefreshTokenUsecase(userRepo user.UserRepo) auth.RefreshTokenUsecase {
	return &RefreshTokenUsecase{
		getByIdAction: &userRepo,
	}
}

// Execute implements auth.RefreshTokenUsecase.
func (uc *RefreshTokenUsecase) Execute(id uint64) *helper.Response {
	res, err := uc.getByIdAction.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.GetResponse("NotFound: data user not found", 500, true)
		}
		return helper.GetResponse(err.Error(), 500, true)
	}

	accessToken, err := helper.NewToken(uint64(res.Id), res.IsAdmin).Create()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	refreshToken, err := helper.NewRefresh(uint64(res.Id), res.IsAdmin).CreateRefresh()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return helper.GetResponse(response, 200, false)
}
