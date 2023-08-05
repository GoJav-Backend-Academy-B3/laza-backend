package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsEmail(email string) bool {
	var data model.User
	err := r.db.
		First(&data, "email = ?", email)
	return err.Error == nil
}
