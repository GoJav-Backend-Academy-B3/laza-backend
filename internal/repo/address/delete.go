package address

import "github.com/phincon-backend/laza/domain/model"

func (r *AddressRepo) Delete(id any) error {
	var address model.Address
	err := r.db.Where("id = ?", id).Delete(&address).Error
	if err != nil {
		return err
	}

	return nil
}
