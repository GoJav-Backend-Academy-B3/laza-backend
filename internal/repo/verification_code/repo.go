package verification_code

import "gorm.io/gorm"

type VerificationCodeRepo struct {
	db *gorm.DB
}

func NewVerificationCodeRepo(db *gorm.DB) *VerificationCodeRepo {
	return &VerificationCodeRepo{db: db}
}
