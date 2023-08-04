package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetWithLimit(limit, offset uint) (*[]model.User, error) {
	var es []model.User
	if err := r.db.Find(&es).Offset(int(offset)).Limit(int(limit)).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &es, nil
}
