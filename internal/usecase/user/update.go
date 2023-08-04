package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type UpdateUserUsecase struct {
	updateAction repositories.UpdateAction[model.User]
}

func NewUpdateUserUsecase(repo repositories.UpdateAction[model.User]) user.UpdateUserUsecase {
	return &UpdateUserUsecase{updateAction: repo}
}

// Excute implements user.UpdateUserUsecase.
func (uc *UpdateUserUsecase) Excute(id uint64, user model.User) *helper.Response {
	result, err := uc.updateAction.Update(id, user)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
