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

func (fb *facebookAuthUsecaseImpl) Execute(fbResponse response.FBAuthResponse) (accessToken string, err error) {
	var userDTO = new(model.User)
	userDTO.Email = fbResponse.Email
	userDTO.Username = helper.ExtractUsernameFromEmail(fbResponse.Email)
	userDTO.FullName = fbResponse.Name
	userDTO.ImageUrl = fbResponse.Picture.Data.URL
	userDTO.IsAdmin = false
	userDTO.IsVerified = true

	user, err := fb.findByEmail.FindByEmail(userDTO.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user, err = fb.insertUserAction.Insert(*userDTO)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	jwt := helper.NewToken(uint64(user.Id), false)

	accessToken, err = jwt.Create()
	if err != nil {
		return "", err
	}
	return
}

func NewFacebookAuthUsecase(userRepo user.UserRepo) facebook_auth.FacebookAuthUsecase {
	return &facebookAuthUsecaseImpl{
		insertUserAction: &userRepo,
		findByEmail:      &userRepo,
	}
}
