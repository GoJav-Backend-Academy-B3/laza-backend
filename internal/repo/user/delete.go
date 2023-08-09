package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) Delete(id any) (err error) {
	tx := r.db.Delete(&model.User{}, "id = ?", id)
	err = tx.Error
	return
}
