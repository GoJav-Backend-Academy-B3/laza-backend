package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) GetById(id any) (e model.User, err error) {
	tx := r.db.First(&e, "id = ?", id)
	err = tx.Error
	return
}
