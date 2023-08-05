package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) Delete(id any) (err error) {
	var user model.User
	tx := r.db.Delete(user, "id = ?", id)
	err = tx.Error
	return
}
