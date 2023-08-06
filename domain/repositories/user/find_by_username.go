package user

import "github.com/phincon-backend/laza/domain/response"

type FindByUsername interface {
	FindByUsername(username string) (response.User, error)
}
