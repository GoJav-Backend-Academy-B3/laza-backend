package facebook_auth

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/facebook_auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"gorm.io/gorm"
)

type facebookAuthUsecaseImpl struct {
	insertUserAction repositories.InsertAction[model.User]
	findByEmail      action.FindByEmail
}

func (fb *facebookAuthUsecaseImpl) Execute(fbResponse response.FBAuthResponse) (accessToken string, refreshToken string, err error) {
	var userDTO = &model.User{
		Email:      fbResponse.Email,
		Username:   helper.ExtractUsernameFromEmail(fbResponse.Email),
		FullName:   fbResponse.Name,
		ImageUrl:   fbResponse.Picture.Data.URL,
		IsAdmin:    false,
		IsVerified: true,
	}

	userByEmail, err := fb.findByEmail.FindByEmail(userDTO.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userByEmail, err = fb.insertUserAction.Insert(*userDTO)
			if err != nil {
				return "", "", err
			}
		} else {
			return "", "", err
		}
	}
	jwt := helper.NewToken(uint64(userByEmail.Id), false)

	accessToken, err = jwt.Create()
	if err != nil {
		return "", "", err
	}
	refreshToken, err = helper.NewRefresh(uint64(userByEmail.Id), userByEmail.IsAdmin).CreateRefresh()
	if err != nil {
		return "", "", err
	}
	return
}

func NewFacebookAuthUsecase(userRepo user.UserRepo) facebook_auth.FacebookAuthUsecase {
	return &facebookAuthUsecaseImpl{
		insertUserAction: &userRepo,
		findByEmail:      &userRepo,
	}
}
