package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsUsername(username string) bool {
	var data model.User
	err := r.db.
		First(&data, "username = ?", username)
	return err.Error == nil
}
