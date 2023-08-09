package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) Count() (c int64, err error) {
	tx := r.db.Model(&model.User{}).Count(&c)
	err = tx.Error
	return
}
