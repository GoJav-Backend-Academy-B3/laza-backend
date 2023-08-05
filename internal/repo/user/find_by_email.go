package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) FindByEmail(email string) (e model.User, err error) {
	tx := r.db.First(&e, "email = ?", email)
	err = tx.Error
	return
}
