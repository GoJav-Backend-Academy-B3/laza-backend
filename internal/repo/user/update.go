package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Update(id uint64, dao model.User) (e model.User, err error) {
	if err := r.db.Model(dao).Where("id = ?", id).Updates(&dao).Error; err != nil {
		return e, errors.New("failed to update data")
	}

	return
}
