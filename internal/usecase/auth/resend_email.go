package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
)

type ResendEmailUserUsecase struct {
	updateAction      repositories.UpdateAction[model.VerificationToken]
	emailAction       actionUser.FindByEmail
	emailExistsAction actionUser.ExistsEmail
}

func NewResendEmailUserUsecase(
	repo repositories.UpdateAction[model.VerificationToken],
	emailExistsAction actionUser.ExistsEmail,
	emailAction actionUser.FindByEmail,
) auth.ResendEmailUserUsecase {
	return &ResendEmailUserUsecase{
		updateAction:      repo,
		emailAction:       emailAction,
		emailExistsAction: emailExistsAction,
	}
}

// Execute implements auth.ResendEmailUserUsecase.
func (uc *ResendEmailUserUsecase) Execute(email string) *helper.Response {
	if emailExists := uc.emailExistsAction.ExistsEmail(email); !emailExists {
		return helper.GetResponse("email is not registered", 500, true)
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

	_, err = uc.updateAction.Update(data.Id, daoToken)
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
