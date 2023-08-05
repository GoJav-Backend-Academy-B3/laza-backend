package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Update(id any, dao model.User) (e model.User, err error) {
	tx := r.db.Model(dao).Where("id = ?", id).Updates(&dao).Scan(&e)
	err = tx.Error
	return
}
