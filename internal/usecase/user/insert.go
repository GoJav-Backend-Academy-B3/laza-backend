package user

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"github.com/phincon-backend/laza/internal/repo/verification_token"
)

type InsertUserUsecase struct {
	insertTokenAction    repositories.InsertAction[model.VerificationToken]
	insertUserAction     repositories.InsertAction[model.User]
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewInsertUserUsecase(userRepo user.UserRepo, tokenRepo verification_token.VerificationTokenRepo) contract.InsertUserUsecase {
	return &InsertUserUsecase{
		insertTokenAction:    &tokenRepo,
		insertUserAction:     &userRepo,
		emailExistsAction:    &userRepo,
		usernameExistsAction: &userRepo,
	}
}

// Excute implements user.InsertUserUsecase.
func (uc *InsertUserUsecase) Execute(user requests.Register) *helper.Response {
	if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
		return helper.GetResponse("username is taken, try another", 500, true)
	}

	if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
		return helper.GetResponse("email is taken, try another", 500, true)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	codeVerify := helper.GenerateRandomNumericString(4)
	expiryDate, _ := helper.GetExpiryDate(5*time.Minute, "Asia/Jakarta")
	daoToken := model.VerificationToken{
		Token:      codeVerify,
		ExpiryDate: expiryDate,
	}

	dao := model.User{
		FullName:           user.FullName,
		Username:           user.Username,
		Password:           hashPassword,
		Email:              user.Email,
		VerificationTokens: []model.VerificationToken{daoToken},
	}

	res, err := uc.insertUserAction.Insert(dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	configMail := helper.DataMail{
		Username: res.Username,
		Email:    res.Email,
		Token:    codeVerify,
		Subject:  "Your verification account",
	}

	err = helper.Mail(&configMail).Send()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	result := response.UserModelResponse(res)
	return helper.GetResponse(result, 201, false)
}
