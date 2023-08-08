package auth

import (
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
)

type LoginGoogleUserUsecase struct {
	insertUserAction     repositories.InsertAction[response.User]
	findByEmailAction    action.FindByEmail
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewLoginGoogleUserUsecase(
	insertUserAction repositories.InsertAction[response.User],
	findByEmailAction action.FindByEmail,
	emailExistsAction action.ExistsEmail,
	usernameExistsAction action.ExistsUsername,
) auth.LoginGoogleUserUsecase {
	return &LoginGoogleUserUsecase{
		insertUserAction:     insertUserAction,
		findByEmailAction:    findByEmailAction,
		emailExistsAction:    emailExistsAction,
		usernameExistsAction: usernameExistsAction,
	}
}

// Excute implements user.InsertUserUsecase.
func (uc *LoginGoogleUserUsecase) Execute(user *helper.GoogleUserResult) *helper.Response {
	username := helper.ExtractUsernameFromEmail(user.Email)
	if userExists := uc.usernameExistsAction.ExistsUsername(username); userExists {
		return helper.GetResponse("username is already registered", 500, true)
	}

	if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); !emailExists {
		data := response.User{
			FullName:   user.Name,
			Username:   username,
			Email:      user.Email,
			ImageUrl:   user.Picture,
			IsVerified: user.Verified_email,
		}

		result, err := uc.insertUserAction.Insert(data)
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}

		jwt := helper.NewToken(uint64(result.Id), result.IsAdmin)

		accessToken, err := jwt.Create()
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}

		response := map[string]string{
			"access_token": accessToken,
		}

		return helper.GetResponse(response, 200, false)

	}

	result, err := uc.findByEmailAction.FindByEmail(user.Email)
	if err != nil {
		return helper.GetResponse("user is not exist", 500, true)
	}

	jwt := helper.NewToken(uint64(result.Id), result.IsAdmin)

	accessToken, err := jwt.Create()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"access_token": accessToken,
	}

	return helper.GetResponse(response, 200, false)
}
