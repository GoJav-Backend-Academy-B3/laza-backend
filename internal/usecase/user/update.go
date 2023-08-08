package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/user"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type UpdateUserUsecase struct {
	updateAction         repositories.UpdateAction[model.User]
	getByIdAction        repositories.GetByIdAction[model.User]
	emailExistsAction    action.ExistsEmail
	usernameExistsAction action.ExistsUsername
}

func NewUpdateUserUsecase(userRepo user.UserRepo) contract.UpdateUserUsecase {
	return &UpdateUserUsecase{
		updateAction:         &userRepo,
		getByIdAction:        &userRepo,
		emailExistsAction:    &userRepo,
		usernameExistsAction: &userRepo,
	}
}

// Excute implements user.UpdateUserUsecase.
func (uc *UpdateUserUsecase) Execute(id uint64, user requests.UpdateUser) *helper.Response {
	data, err := uc.getByIdAction.GetById(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)

	}

	if data.Email != user.Email {
		if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
			return helper.GetResponse("email is already registered", 500, true)
		}
		data.Email = user.Email
		data.IsVerified = false
	}

	if data.Username != user.Username {
		if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
			return helper.GetResponse("username is already registered", 500, true)
		}
		data.Username = user.Username
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
		FullName:   user.FullName,
		Username:   data.Username,
		Email:      data.Email,
		ImageUrl:   imageUrl,
		IsVerified: data.IsVerified,
	}

	res, err := uc.updateAction.Update(id, dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	result := response.UserModelResponse(res)
	return helper.GetResponse(result, 200, false)
}
