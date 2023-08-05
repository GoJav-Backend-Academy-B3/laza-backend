package user

import "github.com/phincon-backend/laza/domain/model"

type FindByEmail interface {
	FindByEmail(email string) (model.User, error)
}