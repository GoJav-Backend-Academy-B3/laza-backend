package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type ChangePasswordUserUsecase struct {
	updateAction  repositories.UpdateAction[model.User]
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewChangePasswordUserUsecase(
	updateAction repositories.UpdateAction[model.User],
	getByIdAction repositories.GetByIdAction[model.User],
) user.ChangePasswordUserUsecase {
	return &ChangePasswordUserUsecase{
		updateAction:  updateAction,
		getByIdAction: getByIdAction,
	}
}

// Execute implements auth.ChangePasswordUserUsecase.
func (uc *ChangePasswordUserUsecase) Execute(id uint64, user requests.ChangePassword) *helper.Response {
	data, err := uc.getByIdAction.GetById(id)
	if err != nil {
		return helper.GetResponse("user is not exist", 500, true)
	}

	if !helper.CheckPassword(data.Password, user.OldPassword) {
		return helper.GetResponse("old password wrong", 500, true)
	}

	if user.NewPassword != user.RePassword {
		return helper.GetResponse("passwords do not match. please try again.", 500, true)
	}

	hashPassword, err := helper.HashPassword(user.NewPassword)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	dao := model.User{
		Password: hashPassword,
	}
	_, err = uc.updateAction.Update(data.Id, dao)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	response := map[string]string{
		"message": "successfully update password",
	}

	return helper.GetResponse(response, 200, false)
}
