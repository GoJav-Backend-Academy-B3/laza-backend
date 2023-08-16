package category

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (cr *CategoryRepo) FindById(id uint64) (category model.Category, err error) {
	db := cr.db.Where("id = ?", id).First(&category)
	err = db.Error
	return
}

func (cr *CategoryRepo) GetById(id_r any) (category model.Category, err error) {
	id := id_r.(uint64)
	db := cr.db.Where("id = ?", id).First(&category)
	err = db.Error
	return
}
