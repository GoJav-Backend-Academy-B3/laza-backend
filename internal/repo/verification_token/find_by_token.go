package verification_token

import "github.com/phincon-backend/laza/domain/model"

func (r *VerificationTokenRepo) FindByToken(id uint64, token string) (e model.VerificationToken, err error) {
	tx := r.db.First(&e, "user_id = ? AND token = ?", id, token)
	err = tx.Error
	return
}
