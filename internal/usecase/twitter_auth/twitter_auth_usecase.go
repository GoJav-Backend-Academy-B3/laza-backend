package twitterauth

import (
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

func (uc *twitterAuthUsecase) Execute(rb response.TwitterUser) (_result any, err error) {
	var userDAO = new(model.User)

	if exist, emailExist := uc.existByUsernameAction.ExistsUsername(rb.NickName), uc.emailExistsAction.ExistsEmail(rb.Email); !exist && !emailExist {

		insertedUser, err := uc.insertUserAction.Insert(*userDAO)
		userDAO.Id, userDAO.IsAdmin = insertedUser.Id, insertedUser.IsAdmin
		if err != nil {
			return _result, err
		}

	} else {
		user, err := uc.findByEmailAction.FindByEmail(rb.Email)
		userDAO.Id, userDAO.IsAdmin = user.Id, user.IsAdmin

		if err != nil {
			return _result, err
		}

	}

	jwt := helper.NewToken(uint64(userDAO.Id), userDAO.IsAdmin)
	accessToken, err := jwt.Create()

	_result = map[string]string{"access_token": accessToken}
	return
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
