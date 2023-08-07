package user

import "github.com/phincon-backend/laza/domain/response"

func (r *UserRepo) FindByUsername(username string) (e response.User, err error) {
	tx := r.db.First(&e, "username = ?", username)
	err = tx.Error
	return
}
