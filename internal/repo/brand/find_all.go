package brand

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *BrandRepo) GetAll() (brands []model.Brand, err error) {
	err = r.db.Find(&brands).Error
	if err != nil {
		return
	}

	return
}
