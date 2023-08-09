package verification_code

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *VerificationCodeRepo) FindByUserId(id any) (e model.VerificationCode, err error) {
	tx := r.db.First(&e, "user_id = ?", id)
	err = tx.Error
	return
}
