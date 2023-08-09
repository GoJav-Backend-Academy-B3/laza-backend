package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
	"gorm.io/gorm"
)

type ChangePasswordUserUsecase struct {
	updateAction  repositories.UpdateAction[model.User]
	getByIdAction repositories.GetByIdAction[model.User]
}

func NewChangePasswordUserUsecase(userRepo user.UserRepo) contract.ChangePasswordUserUsecase {
	return &ChangePasswordUserUsecase{
		updateAction:  &userRepo,
		getByIdAction: &userRepo,
	}
}

// Execute implements auth.ChangePasswordUserUsecase.
func (uc *ChangePasswordUserUsecase) Execute(id uint64, user requests.ChangePassword) *helper.Response {
	data, err := uc.getByIdAction.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return helper.GetResponse("NotFound: data user not found", 500, true)
		}
		return helper.GetResponse(err.Error(), 500, true)
	}

	if !helper.CheckPassword(data.Password, user.OldPassword) {
		return helper.GetResponse("old password is invalid", 500, true)
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
