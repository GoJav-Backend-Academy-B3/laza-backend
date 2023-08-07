package address

import "github.com/phincon-backend/laza/domain/model"

func (r *addressRepo) Update(id any, address model.Address) (model.Address, error) {
	err := r.db.Model(&address).Where("id = ?", id).Updates(&address).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
