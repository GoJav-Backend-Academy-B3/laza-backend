package address

import "github.com/phincon-backend/laza/domain/model"

func (r *AddressRepo) SetAllAddressesNonPrimary(userId uint64) (err error) {
	err = r.db.Model(&model.Address{}).Where("user_id = ?", userId).Update("is_primary", false).Error

	return
}
