package category

import "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) Update(id uint64, model model.Category) (category model.Category, err error) {
	tx := cr.db.First(&model, "id = ?", id)
	err = tx.Error
	if err != nil {
		return
	}

	// Modify data
	category.Update(model)

	// Update data. Please keep in mind that this Save function
	// inserts data if no matching primary key found
	tx = cr.db.Save(&category)
	err = tx.Error

	return
}
