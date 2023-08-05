package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetAll() (es []model.User, err error) {
	tx := r.db.Find(&es)
	err = tx.Error
	return
}
