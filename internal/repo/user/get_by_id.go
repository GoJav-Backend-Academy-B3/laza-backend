package user

import "github.com/phincon-backend/laza/domain/response"

func (r *UserRepo) GetById(id any) (e response.User, err error) {
	tx := r.db.First(&e, "id = ?", id)
	err = tx.Error
	return
}
