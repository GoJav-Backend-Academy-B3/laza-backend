package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetAll() (*[]model.User, error) {
	var es []model.User
	if err := r.db.Find(&es).Error; err != nil {
		return nil, errors.New("failed to get data")
	}
	return &es, nil
}
