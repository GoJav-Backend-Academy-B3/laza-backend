package user

import "github.com/phincon-backend/laza/domain/response"

func (r *UserRepo) FindByEmail(email string) (e response.User, err error) {
	tx := r.db.First(&e, "email = ?", email)
	err = tx.Error
	return
}
