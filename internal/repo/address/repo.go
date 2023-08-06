package address

import (
	"gorm.io/gorm"
)

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *addressRepo {
	return &addressRepo{db: db}
}
