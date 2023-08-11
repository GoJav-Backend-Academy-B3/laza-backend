package address

import "github.com/phincon-backend/laza/domain/model"

func (r *AddressRepo) FindLatestAddressByUserId(user_id uint64) (address model.Address) {
	err := r.db.Where("user_id = ?", user_id).Order("id desc").First(&address).Error
	if err != nil {
		return
	}

	return
}
