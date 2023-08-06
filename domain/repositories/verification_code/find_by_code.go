package verificationcode

import "github.com/phincon-backend/laza/domain/model"

type FindByCode interface {
	FindByCode(id uint64, code string) (model.VerificationCode, error)
}
