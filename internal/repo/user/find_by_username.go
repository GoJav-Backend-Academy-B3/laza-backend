package user

import "github.com/phincon-backend/laza/domain/model"

func (r *UserRepo) FindByUsername(username string) (e model.User, err error) {
	tx := r.db.First(&e, "username = ?", username)
	err = tx.Error
	return
}
