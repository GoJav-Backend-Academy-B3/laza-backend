package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) Delete(id any) (err error) {
	var data model.User
	tx := r.db.Delete(data, "id = ?", id)
	err = tx.Error
	return
}
