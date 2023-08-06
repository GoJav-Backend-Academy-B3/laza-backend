package verificationcode

import "github.com/phincon-backend/laza/domain/model"

func (r *VerificationCodeRepo) FindByCode(id uint64, code string) (e model.VerificationCode, err error) {
	tx := r.db.First(&e, "user_id = ? AND code = ?", id, code)
	err = tx.Error
	return
}
