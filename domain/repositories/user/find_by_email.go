package user

import "github.com/phincon-backend/laza/domain/response"

type FindByEmail interface {
	FindByEmail(email string) (response.User, error)
}
