package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionToken "github.com/phincon-backend/laza/domain/repositories/verification_token"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"github.com/phincon-backend/laza/internal/repo/verification_token"
)

type VerifyEmailUserUsecase struct {
	updateAction repositories.UpdateAction[model.User]
	emailAction  actionUser.FindByEmail
	tokenAction  actionToken.FindByToken
}

func NewVerifyEmailUserUsecase(userRepo user.UserRepo, tokenRepo verification_token.VerificationTokenRepo) auth.VerifyEmailUserUsecase {
	return &VerifyEmailUserUsecase{
		updateAction: &userRepo,
		emailAction:  &userRepo,
		tokenAction:  &tokenRepo,
	}
}

// Execute implements auth.VerifyEmailUserUsecase.
func (uc *VerifyEmailUserUsecase) Execute(email, token string) *helper.Response {
	if email == "" && token == "" {
		return helper.GetResponse("email and token are both empty", 400, true)
	} else if email == "" {
		return helper.GetResponse("email empty", 400, true)
	} else if token == "" {
		return helper.GetResponse("token empty", 400, true)
	}

	dataUser, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse("email is not exist", 500, true)
	}

	dataToken, err := uc.tokenAction.FindByToken(uint64(dataUser.Id), token)
	if err != nil {
		return helper.GetResponse("failed to verify email", 500, true)
	}

	location, _ := time.LoadLocation("Asia/Jakarta")

	if dataUser.IsVerified {
		return helper.GetResponse("already registered, you can login", 500, true)
	} else if dataToken.Token != token {
		return helper.GetResponse("failed to verify email", 500, true)
	} else if dataToken.ExpiryDate.In(location).Add(-7 * time.Hour).Before(time.Now().In(location)) {
		return helper.GetResponse("expired verify email, please resend verify again!", 500, true)
	}

	dao := model.User{
		IsVerified: true,
	}
	_, err = uc.updateAction.Update(dataUser.Id, dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "successfully verification email",
	}

	return helper.GetResponse(response, 200, false)
}
