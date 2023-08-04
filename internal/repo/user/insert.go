package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Insert(dao model.User) (*model.User, error) {
	if err := r.db.Create(dao).Error; err != nil {
		return nil, errors.New("failed to create data")
	}

	return &dao, nil
}
