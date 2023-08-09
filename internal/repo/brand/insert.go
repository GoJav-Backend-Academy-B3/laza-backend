package brand

import "github.com/phincon-backend/laza/domain/model"

func (r *BrandRepo) Insert(e model.Brand) (brand model.Brand, err error) {
	db := r.db.Create(&e)
	err = db.Error
	return e, err
}
