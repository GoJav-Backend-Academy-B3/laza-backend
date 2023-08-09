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
	emailExistsAction     action.ExistsEmail
}

func (uc *twitterAuthUsecase) Execute(rp response.TwitterFieldResponse) *helper.Response {
	var userDAO = new(model.User)
	rp.TwitterUser(userDAO)

	response := map[string]string{}
	if exist, emailExist := uc.existByUsernameAction.ExistsUsername(userDAO.Username), uc.emailExistsAction.ExistsEmail(userDAO.Email); !exist && !emailExist {

		insertedUser, err := uc.insertUserAction.Insert(*userDAO)
		userDAO.Id, userDAO.IsAdmin = insertedUser.Id, insertedUser.IsAdmin

		if err != nil {
			return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
		}
	} else {
		user, err := uc.findByEmailAction.FindByEmail(rp.Email)
		userDAO.Id, userDAO.IsAdmin = user.Id, user.IsAdmin

		if err != nil {
			return helper.GetResponse("user is not exist", 401, true)
		}
	}

	jwt := helper.NewToken(uint64(userDAO.Id), userDAO.IsAdmin)
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
	emailExistsAction action.ExistsEmail,

) tAuth.TwitterAuthUsecase {
	return &twitterAuthUsecase{
		existByUsernameAction: existByUsernameAction,
		insertUserAction:      insertUserAction,
		findByEmailAction:     findByEmailAction,
		existByEmailAction:    existByEmailAction,
		emailExistsAction:     emailExistsAction,
	}
}
