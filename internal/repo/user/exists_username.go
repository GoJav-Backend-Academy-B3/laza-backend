package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsUsername(username string) bool {
	var data model.User
	err := r.db.First(&data, "LOWER(username) = ?", strings.ToLower(username))
	return err.Error == nil
}
