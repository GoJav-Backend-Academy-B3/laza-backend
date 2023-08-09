package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *AddressRepo) GetAllByUserId(userId uint64) (addresses []model.Address, err error) {
	err = r.db.Where("user_id = ?", userId).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, full_name")
	}).Find(&addresses).Error
	if err != nil {
		return
	}

	return
}
