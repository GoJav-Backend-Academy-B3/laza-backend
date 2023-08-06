package verificationtoken

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *VerificationTokenRepo) Insert(dao model.VerificationToken) (e model.VerificationToken, err error) {
	tx := r.db.Create(&dao).Scan(&e)
	err = tx.Error
	return
}
