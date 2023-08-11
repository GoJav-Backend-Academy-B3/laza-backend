package auth

import (
	"time"

	actionUser "github.com/phincon-backend/laza/domain/repositories/user"
	actionCode "github.com/phincon-backend/laza/domain/repositories/verification_code"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"github.com/phincon-backend/laza/internal/repo/verification_code"
)

type VerificationCodeUserUsecase struct {
	emailAction actionUser.FindByEmail
	codeAction  actionCode.FindByCode
}

func NewVerificationCodeUserUsecase(userRepo user.UserRepo, codeRepo verification_code.VerificationCodeRepo) auth.VerificationCodeUserUsecase {
	return &VerificationCodeUserUsecase{
		emailAction: &userRepo,
		codeAction:  &codeRepo,
	}
}

// Execute implements auth.VerificationCodeUserUsecase.
func (uc *VerificationCodeUserUsecase) Execute(user requests.VerificationCode) *helper.Response {
	dataUser, err := uc.emailAction.FindByEmail(user.Email)
	if err != nil {
		return helper.GetResponse("email is invalid", 500, true)
	}

	dataCode, err := uc.codeAction.FindByCode(uint64(dataUser.Id), user.Code)
	if err != nil {
		return helper.GetResponse("code is invalid", 500, true)
	}

	location, _ := time.LoadLocation("Asia/Jakarta")
	if dataCode.Code != user.Code {
		return helper.GetResponse("code is invalid", 500, true)
	} else if dataCode.ExpiryDate.In(location).Add(-7 * time.Hour).Before(time.Now().In(location)) {
		return helper.GetResponse("expired verify email, please resend mail verify again!", 500, true)
	}

	response := map[string]string{
		"message": "code is valid",
	}

	return helper.GetResponse(response, 202, false)
}
