package auth

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type LoginGoogleUserUsecase struct {
	insertUserAction     repositories.InsertAction[model.User]
	findByEmailAction    action.FindByEmail
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewLoginGoogleUserUsecase(userRepo user.UserRepo) auth.LoginGoogleUserUsecase {
	return &LoginGoogleUserUsecase{
		insertUserAction:     &userRepo,
		findByEmailAction:    &userRepo,
		emailExistsAction:    &userRepo,
		usernameExistsAction: &userRepo,
	}
}

// Excute implements user.InsertUserUsecase.
func (uc *LoginGoogleUserUsecase) Execute(user *helper.GoogleUserResult) *helper.Response {
	var data model.User
	if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); !emailExists {
		dao := model.User{
			FullName:   user.Name,
			Username:   helper.ExtractUsernameFromEmail(user.Email),
			Email:      user.Email,
			ImageUrl:   user.Picture,
			IsVerified: user.Verified_email,
		}

		result, err := uc.insertUserAction.Insert(dao)
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}

		data = result
	} else {
		result, err := uc.findByEmailAction.FindByEmail(user.Email)
		if err != nil {
			return helper.GetResponse("user is not exist", 500, true)
		}

		data = result
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
