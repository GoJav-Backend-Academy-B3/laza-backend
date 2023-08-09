package verification_code

import "github.com/phincon-backend/laza/domain/model"

func (r *VerificationCodeRepo) Update(id any, dao model.VerificationCode) (e model.VerificationCode, err error) {
	tx := r.db.Model(dao).Where("user_id = ?", id).Updates(&dao).Scan(&e)
	err = tx.Error
	return
}
