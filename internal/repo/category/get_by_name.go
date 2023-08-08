package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) GetByName(category string) (m model.Category, err error) {
	db := cr.db.Where("category = ?", category).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}
