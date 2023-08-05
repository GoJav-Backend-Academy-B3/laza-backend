package user

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase struct {
	insertUserAction     repositories.InsertAction[model.User]
	insertTokenAction    repositories.InsertAction[model.VerificationToken]
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewInsertUserUsecase(repoUser repositories.InsertAction[model.User],
	repoToken repositories.InsertAction[model.VerificationToken], repoExistsEmail action.ExistsEmail,
	repoExistsUsername action.ExistsUsername) user.InsertUserUsecase {
	return &InsertUserUsecase{
		insertUserAction:     repoUser,
		insertTokenAction:    repoToken,
		emailExistsAction:    repoExistsEmail,
		usernameExistsAction: repoExistsUsername,
	}
}

// Excute implements user.InsertUserUsecase.
func (uc *InsertUserUsecase) Execute(user model.User) *helper.Response {
	if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
		return helper.GetResponse("username is already registered", 401, true)
	}

	if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
		return helper.GetResponse("email is already registered", 401, true)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	user.Password = hashPassword

	result, err := uc.insertUserAction.Insert(user)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	codeVerify := helper.GenerateRandomNumericString(4)
	expiryDate, _ := helper.GetExpiryDate(5*time.Minute, "Asia/Jakarta")
	daoToken := model.VerificationToken{
		Token:      codeVerify,
		ExpiryDate: expiryDate,
		UserId:     uint64(result.Id),
	}

	_, err = uc.insertTokenAction.Insert(daoToken)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	configMail := helper.DataMail{
		Username: result.Username,
		Email:    result.Email,
		Token:    codeVerify,
		Subject:  "Verify your email",
	}

	err = helper.Mail(&configMail).Send()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
