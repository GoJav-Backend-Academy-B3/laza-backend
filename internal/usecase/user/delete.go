package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	contract "github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
	"github.com/phincon-backend/laza/internal/repo/user"
)

type DeleteUserUsecase struct {
	deleteAction repositories.DeleteAction[model.User]
}

func NewDeleteUserUsecase(userRepo user.UserRepo) contract.DeleteUserUsecase {
	return &DeleteUserUsecase{
		deleteAction: &userRepo,
	}
}

// Excute implements user.DeleteUserUsecase.
func (uc *DeleteUserUsecase) Execute(id uint64) *helper.Response {
	err := uc.deleteAction.Delete(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse("successfully deleted data user", 200, false)
}
