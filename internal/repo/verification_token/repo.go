package verification_token

import (
	"gorm.io/gorm"
)

type VerificationTokenRepo struct {
	db *gorm.DB
}

func NewVerificationTokenRepo(db *gorm.DB) *VerificationTokenRepo {
	return &VerificationTokenRepo{db: db}
}
