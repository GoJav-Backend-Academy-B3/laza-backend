package verificationtoken

import "github.com/phincon-backend/laza/domain/model"

type FindByToken interface {
	FindByToken(id, token string) (model.VerificationToken, error)
}
