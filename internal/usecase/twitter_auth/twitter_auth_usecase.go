package twitterauth

import (
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/response"
	tAuth "github.com/phincon-backend/laza/domain/usecases/twitter_auth"
	"github.com/phincon-backend/laza/helper"
)

type twitterAuthUsecase struct {
	existByUsernameAction action.ExistsUsername
	insertUserAction      repositories.InsertAction[response.User]
}

func (uc *twitterAuthUsecase) Execute(response.TwitterResponse) *helper.Response {
	return nil
}

func NewtwitterAuthUsecase(existByUsernameAction action.ExistsUsername, insertUserAction repositories.InsertAction[response.User]) tAuth.TwitterAuthUsecase {
	return &twitterAuthUsecase{
		existByUsernameAction: existByUsernameAction,
		insertUserAction:      insertUserAction,
	}
}
