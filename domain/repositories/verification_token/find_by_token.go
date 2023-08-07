package verificationtoken

import "github.com/phincon-backend/laza/domain/model"

type FindByToken interface {
	FindByToken(id uint64, token string) (model.VerificationToken, error)
}
