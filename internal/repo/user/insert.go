package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) Insert(dao model.User) (e model.User, err error) {
	tx := r.db.Create(&dao).Scan(&e)
	err = tx.Error
	return
}
