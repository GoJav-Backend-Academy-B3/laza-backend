package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) ExistsEmail(email string) bool {
	err := r.db.First(&model.User{}, "LOWER(email) = ?", strings.ToLower(email))
	return err.Error == nil
}
