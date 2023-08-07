package auth

import (
	"time"

	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionCode "github.com/phincon-backend/laza/domain/repositories/verification_code"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
)

type UpdatePasswordUserUsecase struct {
	updateAction repositories.UpdateAction[response.User]
	emailAction  actionUser.FindByEmail
	codeAction   actionCode.FindByCode
}

func NewUpdatePasswordUserUsecase(
	repo repositories.UpdateAction[response.User],
	emailAction actionUser.FindByEmail,
	codeAction actionCode.FindByCode,
) auth.UpdatePasswordUserUsecase {
	return &UpdatePasswordUserUsecase{
		updateAction: repo,
		emailAction:  emailAction,
		codeAction:   codeAction,
	}
}

// Execute implements auth.UpdatePasswordUserUsecase.
func (uc *UpdatePasswordUserUsecase) Execute(email, code string, user request.UpdatePassword) *helper.Response {
	if user.Password != user.RePassword {
		return helper.GetResponse("passwords do not match. please try again.", 401, true)
	}

	dataUser, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse("email is not exist", 401, true)
	}

	dataCode, err := uc.codeAction.FindByCode(uint64(dataUser.Id), code)
	if err != nil {
		return helper.GetResponse("failed to verify email", 401, true)
	}

	location, _ := time.LoadLocation("Asia/Jakarta")

	if dataCode.Code != code {
		return helper.GetResponse("failed to verify email", 401, true)
	} else if dataCode.ExpiryDate.In(location).Add(-7 * time.Hour).Before(time.Now().In(location)) {
		return helper.GetResponse("expired verify email, please resend verify again!", 401, true)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	data := response.User{
		Password: hashPassword,
	}
	_, err = uc.updateAction.Update(dataUser.Id, data)
	if err != nil {
		return helper.GetResponse(err.Error(), 401, true)
	}

	response := map[string]string{
		"message": "successfully update password",
	}

	return helper.GetResponse(response, 200, false)
}
