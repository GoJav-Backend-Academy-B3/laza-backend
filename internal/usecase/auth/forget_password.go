package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionCode "github.com/phincon-backend/laza/domain/repositories/verification_code"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type ForgetPasswordUserUsecase struct {
	updateAction       repositories.UpdateAction[model.VerificationCode]
	insertAction       repositories.InsertAction[model.VerificationCode]
	findByUserIdAction actionCode.FindByUserId
	emailAction        actionUser.FindByEmail
	emailExistsAction  actionUser.ExistsEmail
}

func NewForgetPasswordUserUsecase(
	repo repositories.UpdateAction[model.VerificationCode],
	insertAction repositories.InsertAction[model.VerificationCode],
	findByUserIdAction actionCode.FindByUserId,
	emailExistsAction actionUser.ExistsEmail,
	emailAction actionUser.FindByEmail,
) auth.ForgetPasswordUserUsecase {
	return &ForgetPasswordUserUsecase{
		updateAction:       repo,
		insertAction:       insertAction,
		findByUserIdAction: findByUserIdAction,
		emailAction:        emailAction,
		emailExistsAction:  emailExistsAction,
	}
}

// Execute implements auth.ForgetPasswordUserUsecase.
func (uc *ForgetPasswordUserUsecase) Execute(email string) *helper.Response {
	if emailExists := uc.emailExistsAction.ExistsEmail(email); !emailExists {
		return helper.GetResponse("email is not registered", 401, true)
	}

	data, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	codeVerify := helper.GenerateRandomNumericString(4)
	expiryDate, _ := helper.GetExpiryDate(5*time.Minute, "Asia/Jakarta")
	daoCode := model.VerificationCode{
		Code:       codeVerify,
		ExpiryDate: expiryDate,
		UserId:     uint64(data.Id),
	}

	_, err = uc.findByUserIdAction.FindByUserId(data.Id)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return helper.GetResponse(err.Error(), 500, true)
		}
		_, err = uc.insertAction.Insert(daoCode)
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}
	} else {
		_, err = uc.updateAction.Update(data.Id, daoCode)
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}
	}

	configMail := helper.DataMail{
		Username: data.Username,
		Email:    data.Email,
		Code:     codeVerify,
		Subject:  "Verification Code",
	}

	err = helper.Mail(&configMail).Send()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "successfully send mail forget password",
	}

	return helper.GetResponse(response, 200, false)
}
