package user

import "github.com/phincon-backend/laza/helper"

type GetAllUserUsecase interface {
	Excute() *helper.Response
}
