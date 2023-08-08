package category

import "github.com/phincon-backend/laza/domain/model"

func (r *CategoryRepo) GetByName(category string) (m model.Category, err error) {
	db := r.db.Where("category = ?", category).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}
