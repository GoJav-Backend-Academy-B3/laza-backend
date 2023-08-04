package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Update(id string, dao model.User) (*model.User, error) {
	if err := r.db.Model(dao).Where("id = ?", id).Updates(&dao).Error; err != nil {
		return nil, errors.New("failed to update data")
	}

	return &dao, nil
}
