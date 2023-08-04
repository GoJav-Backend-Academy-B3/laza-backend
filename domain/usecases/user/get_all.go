package user

import "github.com/phincon-backend/laza/helper"

type GetAllUserUsecase interface {
	Execute() *helper.Response
}
