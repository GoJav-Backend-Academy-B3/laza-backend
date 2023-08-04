package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetById(id uint64) (*model.User, error) {
	var e model.User
	if err := r.db.Find(&e, "id = ?", id).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &e, nil
}
