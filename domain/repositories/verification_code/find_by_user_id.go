package verificationcode

import "github.com/phincon-backend/laza/domain/model"

type FindByUserId interface {
	FindByUserId(id any) (model.VerificationCode, error)
}
