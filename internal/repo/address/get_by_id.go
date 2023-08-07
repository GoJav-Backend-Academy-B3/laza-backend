package address

import "github.com/phincon-backend/laza/domain/model"

func (r *addressRepo) GetById(id any) (model.Address, error) {
	var address model.Address

	err := r.db.Preload("users").First(&address, "id = ?", id).Error
	if err != nil {
		return address, err
	}

	return address, nil
}
