package user

import "github.com/phincon-backend/laza/domain/response"

func (r *UserRepo) Update(id any, dao response.User) (e response.User, err error) {
	tx := r.db.Model(dao).Where("id = ?", id).Updates(&dao).Scan(&e)
	err = tx.Error
	return
}
