package user

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *UserRepo) GetById(id any) (e model.User, err error) {
	tx := r.db.First(&e, "id = ?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		err = errors.New("data user not found")
	} else {
		err = tx.Error
	}
	return
}
