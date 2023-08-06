package facebook_auth

import (
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/facebook_auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/helpers"
)

type facebookAuthUsecaseImpl struct {
	existByUsernameAction action.ExistsUsername
	insertUserAction      repositories.InsertAction[response.User]
}

func (fb *facebookAuthUsecaseImpl) Execute(fbResponse response.FBAuthResponse) (accessToken string, err error) {
	var userDAO = new(response.User)
	userDAO.Email = fbResponse.Email
	userDAO.Username = helpers.ExtractUsernameFromEmail(fbResponse.Email)
	userDAO.FullName = fbResponse.Name
	userDAO.ImageUrl = fbResponse.Picture.Data.URL
	userDAO.IsAdmin = false
	userDAO.IsVerified = true

	var insertedUser response.User
	if exist := fb.existByUsernameAction.ExistsUsername(userDAO.Username); !exist {
		insertedUser, err = fb.insertUserAction.Insert(*userDAO)
		if err != nil {
			return "", err
		}
	}

	jwt := helper.NewToken(uint64(insertedUser.Id), false)

	accessToken, err = jwt.Create()
	if err != nil {
		return "", err
	}
	return
}

func NewFacebookAuthUsecase(existByusernameAction action.ExistsUsername) facebook_auth.FacebookAuthUsecase {
	return &facebookAuthUsecaseImpl{
		existByUsernameAction: existByusernameAction,
	}
}
