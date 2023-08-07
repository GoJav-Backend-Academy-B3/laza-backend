package auth

import (

	"github.com/phincon-backend/laza/domain/request"

	"github.com/phincon-backend/laza/domain/requests"

	"github.com/phincon-backend/laza/helper"
)

type LoginUserUsecase interface {

	Execute(user request.Login) *helper.Response

	Execute(user requests.Login) *helper.Response

}
