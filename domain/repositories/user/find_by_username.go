package user

import "github.com/phincon-backend/laza/domain/model"

type FindByUsername interface {
	FindByUsername(username string) (model.User, error)
}
