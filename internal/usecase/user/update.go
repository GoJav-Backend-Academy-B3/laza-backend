package user

import (
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/request"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type UpdateUserUsecase struct {
	updateAction         repositories.UpdateAction[response.User]
	emailActon           action.FindByEmail
	usernameActon        action.FindByUsername
}

func NewUpdateUserUsecase(repo repositories.UpdateAction[response.User],
	emailActon action.FindByEmail, usernameActon action.FindByUsername) user.UpdateUserUsecase {
	return &UpdateUserUsecase{
		updateAction:         repo,
		emailActon:           emailActon,
		usernameActon:        usernameActon,
	}
}

// Excute implements user.UpdateUserUsecase.
func (uc *UpdateUserUsecase) Execute(id uint64, user request.User) *helper.Response {
	dataEmail, err := uc.emailActon.FindByEmail(user.Email)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return helper.GetResponse(err.Error(), 500, true)
		}
		dataEmail.Email = user.Email
		dataEmail.IsVerified = false
	}

	dataUsername, err := uc.usernameActon.FindByUsername(user.Username)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return helper.GetResponse(err.Error(), 500, true)
		}
		dataUsername.Username = user.Username
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	data := response.User{
		FullName:   user.FullName,
		Username:   dataUsername.Username,
		Password:   hashPassword,
		Email:      dataEmail.Email,
		ImageUrl:   user.Image,
		IsVerified: dataEmail.IsVerified,
	}

	result, err := uc.updateAction.Update(id, data)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
