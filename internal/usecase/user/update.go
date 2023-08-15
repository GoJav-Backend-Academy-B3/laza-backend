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

	if data.Username != user.Username || data.Email != user.Email {
		if data.Username != user.Username {
			if userExists := uc.usernameExistsAction.ExistsUsername(user.Username); userExists {
				return helper.GetResponse("username is taken, try another", 500, true)
			}
			data.Username = user.Username
		}

		if data.Email != user.Email {
			if emailExists := uc.emailExistsAction.ExistsEmail(user.Email); emailExists {
				return helper.GetResponse("email is taken, try another", 500, true)
			}
			data.Email = user.Email
			data.IsVerified = false
		}
	}

	var imageUrl = data.ImageUrl
	if user.Image != nil {
		if user.Image.Size >= int64(2*1024*1024) {
			return helper.GetResponse("file size to large", 400, true)
		}

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

	// TODO: Fix bug before update is_verified false and after update back to true
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
