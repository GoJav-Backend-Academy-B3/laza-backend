package user

import (
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase struct {
	updateAction         repositories.UpdateAction[response.User]
	getByIdAction        repositories.GetByIdAction[response.User]
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewUpdateUserUsecase(repo repositories.UpdateAction[response.User], getByIdAction repositories.GetByIdAction[response.User],
	emailExistsAction action.ExistsEmail, usernameExistsAction action.ExistsUsername) user.UpdateUserUsecase {
	return &UpdateUserUsecase{
		updateAction:         repo,
		getByIdAction:        getByIdAction,
		emailExistsAction:    emailExistsAction,
		usernameExistsAction: usernameExistsAction,
	}
}

// Excute implements user.UpdateUserUsecase.
func (uc *UpdateUserUsecase) Execute(id uint64, user requests.User) *helper.Response {
	dataUser, err := uc.getByIdAction.GetById(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)

	}

	if dataUser.Email != user.Email {
		if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
			return helper.GetResponse("email is already registered", 401, true)
		}
		dataUser.Email = user.Email
		dataUser.IsVerified = false
	}

	if dataUser.Username != user.Username {
		if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
			return helper.GetResponse("username is already registered", 401, true)
		}
		dataUser.Username = user.Username
	}

	hashPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	data := response.User{
		FullName:   user.FullName,
		Username:   dataUser.Username,
		Password:   hashPassword,
		Email:      dataUser.Email,
		ImageUrl:   user.Image,
		IsVerified: dataUser.IsVerified,
	}

	result, err := uc.updateAction.Update(id, data)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, false)
}
