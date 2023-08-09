package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsUsername(username string) bool {
	err := r.db.First(&model.User{}, "LOWER(username) = ?", strings.ToLower(username))
	return err.Error == nil
}
