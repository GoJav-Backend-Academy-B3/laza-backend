package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"github.com/phincon-backend/laza/internal/repo/verification_token"
)

type ResendEmailUserUsecase struct {
	emailAction       actionUser.FindByEmail
	emailExistsAction actionUser.ExistsEmail
	updateTokenAction repositories.UpdateAction[model.VerificationToken]
}

func NewResendEmailUserUsecase(userRepo user.UserRepo, tokenRepo verification_token.VerificationTokenRepo) auth.ResendEmailUserUsecase {
	return &ResendEmailUserUsecase{
		emailAction:       &userRepo,
		emailExistsAction: &userRepo,
		updateTokenAction: &tokenRepo,
	}
}

// Execute implements auth.ResendEmailUserUsecase.
func (uc *ResendEmailUserUsecase) Execute(email string) *helper.Response {
	if emailExists := uc.emailExistsAction.ExistsEmail(email); !emailExists {
		return helper.GetResponse("please enter a valid email address", 500, true)
	}

	data, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	codeVerify := helper.GenerateRandomNumericString(4)
	expiryDate, _ := helper.GetExpiryDate(5*time.Minute, "Asia/Jakarta")
	daoToken := model.VerificationToken{
		Token:      codeVerify,
		ExpiryDate: expiryDate,
		UserId:     uint64(data.Id),
	}

	_, err = uc.updateTokenAction.Update(data.Id, daoToken)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	configMail := helper.DataMail{
		Username: data.Username,
		Email:    data.Email,
		Token:    codeVerify,
		Subject:  "Your verification account",
	}

	err = helper.Mail(&configMail).Send()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "successfully resend email",
	}

	return helper.GetResponse(response, 200, false)
}
