package twitterauth

import (
	rp "github.com/phincon-backend/laza/domain/response"
	"github.com/phincon-backend/laza/helper"
)

type TwitterAuthUsecase interface {
	Execute(rp.TwitterFieldResponse) *helper.Response
}
