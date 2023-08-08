package category

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (cr *CategoryRepo) FindById(id uint64) (category model.Category, err error) {
	db := cr.db.Where("id = ?", id).First(&category)
	err = db.Error
	return
}
