package user

import (
	"strings"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) FindByUsername(username string) (e model.User, err error) {
	tx := r.db.First(&e, "LOWER(username) = ?", strings.ToLower(username))
	err = tx.Error
	return
}
