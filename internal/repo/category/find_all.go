package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) FindAll() (categories []model.Category, err error) {
	db := cr.db.Find(&categories)
	err = db.Error
	return
}
