package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) FindByEmail(email string) (e model.User, err error) {
	tx := r.db.First(&e, "LOWER(email) = ?", strings.ToLower(email))
	err = tx.Error
	return
}
