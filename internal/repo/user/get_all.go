package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetAll() (es []model.User, err error) {
	if err := r.db.Find(&es).Error; err != nil {
		return nil, errors.New("failed to get data")
	}
	return
}
