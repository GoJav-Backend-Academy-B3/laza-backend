package twitterauth

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	tAuth "github.com/phincon-backend/laza/domain/usecases/twitter_auth"
	"github.com/phincon-backend/laza/helper"
)

type twitterAuthUsecase struct {
	existByUsernameAction action.ExistsUsername
	existByEmailAction    action.ExistsEmail
	insertUserAction      repositories.InsertAction[model.User]
	findByEmailAction     action.FindByEmail
}

func (uc *twitterAuthUsecase) Execute(rp response.TwitterFieldResponse) *helper.Response {
	var userDAO = new(model.User)
	response := map[string]string{}

	if exist := uc.existByUsernameAction.ExistsUsername(rp.NickName); !exist {
		userDAO.FullName = rp.Name
		userDAO.Email = rp.Name
		userDAO.Username = rp.NickName
		userDAO.ImageUrl = rp.ImageUrl
		userDAO.IsAdmin = false
		userDAO.IsVerified = true

		insertedUser, err := uc.insertUserAction.Insert(*userDAO)
		if err != nil {
			return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		}

		jwt := helper.NewToken(uint64(insertedUser.Id), false)

		accessToken, err := jwt.Create()
		if err != nil {
			return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		}
		response["access_token"] = accessToken

		return helper.GetResponse(response, http.StatusOK, false)
	}

	user, err := uc.findByEmailAction.FindByEmail(rp.Email)

	if err != nil {
		return helper.GetResponse("user is not exist", 401, true)
	}

	jwt := helper.NewToken(uint64(user.Id), user.IsAdmin)

	accessToken, err := jwt.Create()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response["access_token"] = accessToken

	return helper.GetResponse(response, 200, false)

}

func NewtwitterAuthUsecase(
	existByUsernameAction action.ExistsUsername,
	insertUserAction repositories.InsertAction[model.User],
	findByEmailAction action.FindByEmail,
	existByEmailAction action.ExistsEmail,

) tAuth.TwitterAuthUsecase {
	return &twitterAuthUsecase{
		existByUsernameAction: existByUsernameAction,
		insertUserAction:      insertUserAction,
		findByEmailAction:     findByEmailAction,
		existByEmailAction:    existByEmailAction,
	}
}
