package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Insert(dao model.User) (e model.User, err error) {
	if err := r.db.Create(dao).Error; err != nil {
		return e, errors.New("failed to create data")
	}

	return
}
