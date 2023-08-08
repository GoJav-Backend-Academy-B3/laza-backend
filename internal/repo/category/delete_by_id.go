package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) DeleteById(id uint64) (rowAffected int64, err error) {
	db := cr.db.Delete(&model.Category{}, id)
	err = db.Error
	rowAffected = db.RowsAffected
	return
}
