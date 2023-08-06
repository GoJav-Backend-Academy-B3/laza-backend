package user

import (
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type DeleteUserUsecase struct {
	deleteAction repositories.DeleteAction[response.User]
}

func NewDeleteUserUsecase(repo repositories.DeleteAction[response.User]) user.DeleteUserUsecase {
	return &DeleteUserUsecase{deleteAction: repo}
}

// Excute implements user.DeleteUserUsecase.
func (uc *DeleteUserUsecase) Execute(id uint64) *helper.Response {
	err := uc.deleteAction.Delete(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse("successfully deleted data user", 200, false)
}
