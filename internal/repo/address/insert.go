package address

import "github.com/phincon-backend/laza/domain/model"

func (r *AddressRepo) Insert(address model.Address) (model.Address, error) {
	err := r.db.Create(&address).Error
	if err != nil {
		return address, nil
	}

	return address, nil
}
