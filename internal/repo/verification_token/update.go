package verification_token

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *VerificationTokenRepo) Update(id any, dao model.VerificationToken) (e model.VerificationToken, err error) {
	tx := r.db.Model(dao).Where("user_id = ?", id).Updates(&dao).Scan(&e)
	err = tx.Error
	return
}
