package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) Insert(e model.Category) (category model.Category, err error) {
	db := cr.db.Create(&e)
	err = db.Error
	return
}
