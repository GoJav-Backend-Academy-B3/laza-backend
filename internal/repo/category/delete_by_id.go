package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) DeleteById(id uint64) (err error) {
	db := cr.db.Delete(&model.Category{}, id)
	err = db.Error
	return
}
