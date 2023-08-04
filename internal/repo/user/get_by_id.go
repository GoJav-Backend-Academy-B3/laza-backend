package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetById(id uint64) (e model.User, err error) {
	if err := r.db.Find(&e, "id = ?", id).Error; err != nil {
		return e, errors.New("failed to get data")
	}

	return e, nil
}
