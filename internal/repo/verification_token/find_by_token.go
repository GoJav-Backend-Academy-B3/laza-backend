package verificationtoken

import "github.com/phincon-backend/laza/domain/model"

func (r *VerificationTokenRepo) FindByToken(id, token string) (e model.VerificationToken, err error) {
	tx := r.db.First(&e, "id = ? AND token = ?", id, token)
	err = tx.Error
	return
}
