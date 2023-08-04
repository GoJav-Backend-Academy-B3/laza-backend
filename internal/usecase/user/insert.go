package user

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/user"
	"github.com/phincon-backend/laza/helper"
)

type InsertUserUsecase struct {
	insertAction repositories.InsertAction[model.User]
}

func NewInsertUserUsecase(repo repositories.InsertAction[model.User]) user.InsertUserUsecase {
	return &InsertUserUsecase{insertAction: repo}
}

// Excute implements user.InsertUserUsecase.
func (uc *InsertUserUsecase) Execute(user model.User) *helper.Response {
	result, err := uc.insertAction.Insert(user)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
