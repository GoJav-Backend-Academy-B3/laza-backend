package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionToken "github.com/phincon-backend/laza/domain/repositories/verification_token"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
)

type VerifyEmailUserUsecase struct {
	updateAction repositories.UpdateAction[response.User]
	emailAction  actionUser.FindByEmail
	tokenAction  actionToken.FindByToken
}

func NewVerifyEmailUserUsecase(repo repositories.UpdateAction[response.User],
	emailAction actionUser.FindByEmail,
	tokenAction actionToken.FindByToken) auth.VerifyEmailUserUsecase {
	return &VerifyEmailUserUsecase{
		updateAction: repo,
		emailAction:  emailAction,
		tokenAction:  tokenAction,
	}
}

// Execute implements auth.VerifyEmailUserUsecase.
func (uc *VerifyEmailUserUsecase) Execute(email, token string) *helper.Response {
	user, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse("email is not exist", 401, true)
	}

	dataToken, err := uc.tokenAction.FindByToken(uint64(user.Id), token)
	if err != nil {
		return helper.GetResponse("failed to verify email", 401, true)
	}

	location, _ := time.LoadLocation("Asia/Jakarta")

	if user.IsVerified {
		return helper.GetResponse("already registered, you can login", 401, true)
	} else if dataToken.Token != token {
		return helper.GetResponse("failed to verify email", 401, true)
	} else if dataToken.ExpiryDate.In(location).Add(-7 * time.Hour).Before(time.Now().In(location)) {
		return helper.GetResponse("expired verify email, please resend verify again!", 401, true)
	}

	data := response.User{
		IsVerified: true,
	}
	_, err = uc.updateAction.Update(user.Id, data)
	if err != nil {
		return helper.GetResponse(err.Error(), 401, true)
	}

	response := map[string]string{
		"message": "successfully verification email",
	}

	return helper.GetResponse(response, 200, false)
}
