package user

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase struct {
	insertUserAction     repositories.InsertAction[model.User]
	insertTokenAction    repositories.InsertAction[model.VerificationToken]
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewInsertUserUsecase(
	repoUser repositories.InsertAction[model.User],
	repoToken repositories.InsertAction[model.VerificationToken],
	repoExistsEmail action.ExistsEmail,
	repoExistsUsername action.ExistsUsername,
) user.InsertUserUsecase {
	return &InsertUserUsecase{
		insertUserAction:     repoUser,
		insertTokenAction:    repoToken,
		emailExistsAction:    repoExistsEmail,
		usernameExistsAction: repoExistsUsername,
	}
}

// Excute implements user.InsertUserUsecase.
func (uc *InsertUserUsecase) Execute(user requests.Register) *helper.Response {
	if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
		return helper.GetResponse("username is already registered", 500, true)
	}

	if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
		return helper.GetResponse("email is already registered", 500, true)
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	var imageUrl = helper.DefaultImageProfileUrl
	if user.Image != nil {
		file, err := user.Image.Open()
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}
		defer file.Close()

		url, err := helper.UploadImageFile("user", file)
		if err != nil {
			return helper.GetResponse(err.Error(), 500, true)
		}
		imageUrl = url
	}

	dao := model.User{
		FullName: user.FullName,
		Username: user.Username,
		Password: hashPassword,
		Email:    user.Email,
		ImageUrl: imageUrl,
	}

	res, err := uc.insertUserAction.Insert(dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	codeVerify := helper.GenerateRandomNumericString(4)
	expiryDate, _ := helper.GetExpiryDate(5*time.Minute, "Asia/Jakarta")
	daoToken := model.VerificationToken{
		Token:      codeVerify,
		ExpiryDate: expiryDate,
		UserId:     uint64(res.Id),
	}

	_, err = uc.insertTokenAction.Insert(daoToken)
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
