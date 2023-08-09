package verification_code

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *VerificationCodeRepo) Insert(dao model.VerificationCode) (e model.VerificationCode, err error) {
	tx := r.db.Create(&dao).Scan(&e)
	err = tx.Error
	return
}