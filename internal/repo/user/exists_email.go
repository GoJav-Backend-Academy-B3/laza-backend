package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsEmail(email string) bool {
	var data model.User
	err := r.db.First(&data, "LOWER(email) = ?", strings.ToLower(email))
	return err.Error == nil
}
