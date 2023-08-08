package gopay

import (
	"gorm.io/gorm"
)

type GopayRepo struct {
	db *gorm.DB
}

func NewGopayRepo(db *gorm.DB) *GopayRepo {
	return &GopayRepo{db: db}
}
