package address

import "github.com/phincon-backend/laza/domain/model"

func (r *addressRepo) GetAllByUserId(userId uint64) (addresses []model.Address, err error) {
	err = r.db.Where("user_id = ?", userId).Find(&addresses).Error
	if err != nil {
		return
	}

	return
}
