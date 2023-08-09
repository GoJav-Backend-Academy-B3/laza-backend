package auth

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionCode "github.com/phincon-backend/laza/domain/repositories/verification_code"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"github.com/phincon-backend/laza/internal/repo/verification_code"
)

type ResetPasswordUserUsecase struct {
	updateAction repositories.UpdateAction[model.User]
	emailAction  actionUser.FindByEmail
	codeAction   actionCode.FindByCode
}

func NewResetPasswordUserUsecase(userRepo user.UserRepo, codeRepo verification_code.VerificationCodeRepo) auth.ResetPasswordUserUsecase {
	return &ResetPasswordUserUsecase{
		updateAction: &userRepo,
		emailAction:  &userRepo,
		codeAction:   &codeRepo,
	}
}

// Execute implements auth.UpdatePasswordUserUsecase.
func (uc *ResetPasswordUserUsecase) Execute(email string, user requests.ResetPassword) *helper.Response {
	if email == "" {
		return helper.GetResponse("email empty", 400, true)
	}

	if user.NewPassword != user.RePassword {
		return helper.GetResponse("passwords do not match, please try again.", 401, true)
	}

	dataUser, err := uc.emailAction.FindByEmail(email)
	if err != nil {
		return helper.GetResponse("email is invalid", 401, true)
	}

	hashPassword, err := helper.HashPassword(user.NewPassword)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	dao := model.User{
		Password: hashPassword,
	}
	_, err = uc.updateAction.Update(dataUser.Id, dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 401, true)
	}

	response := map[string]string{
		"message": "successfully reset password",
	}

	return helper.GetResponse(response, 200, false)
}
