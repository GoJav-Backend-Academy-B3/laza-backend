package size

import "gorm.io/gorm"

type SizeRepo struct {
	db *gorm.DB
}

func NewSizeRepo(db *gorm.DB) *SizeRepo {
	return &SizeRepo{
		db: db,
	}
}
