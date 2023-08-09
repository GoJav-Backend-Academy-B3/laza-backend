package auth

import (
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type LoginUserUsecase struct {
	usernameActon action.FindByUsername
}

func NewLoginUserUsecase(userRepo user.UserRepo) auth.LoginUserUsecase {
	return &LoginUserUsecase{
		usernameActon: &userRepo,
	}
}

// Execute implements auth.LoginUserUsecase.
func (uc *LoginUserUsecase) Execute(user requests.Login) *helper.Response {
	data, err := uc.usernameActon.FindByUsername(user.Username)
	if err != nil {
		return helper.GetResponse("username or password is invalid", 500, true)
	}

	if !helper.CheckPassword(data.Password, user.Password) {
		return helper.GetResponse("username or password is invalid", 500, true)
	}

	if !data.IsVerified {
		return helper.GetResponse("please verify your account", 500, true)
	}

	accessToken, err := helper.NewToken(uint64(data.Id), data.IsAdmin).Create()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"access_token": accessToken,
	}

	return helper.GetResponse(response, 200, false)
}
