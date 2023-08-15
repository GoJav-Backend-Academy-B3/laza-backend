package twitterauth

import (
	"github.com/phincon-backend/laza/domain/response"
)

type TwitterAuthUsecase interface {
	Execute(rb response.TwitterUser) (_result any, err error)
}
