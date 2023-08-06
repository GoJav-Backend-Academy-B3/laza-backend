package user

import (
	"github.com/phincon-backend/laza/domain/response"
)

func (r *UserRepo) Insert(dao response.User) (e response.User, err error) {
	tx := r.db.Create(&dao).Scan(&e)
	err = tx.Error
	return
}
