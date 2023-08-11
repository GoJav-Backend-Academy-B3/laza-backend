package auth

import "github.com/phincon-backend/laza/helper"

type RefreshTokenUsecase interface {
	Execute(id uint64) *helper.Response
}
