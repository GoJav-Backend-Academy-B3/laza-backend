package category

import model "github.com/phincon-backend/laza/domain/model"

func (cr *CategoryRepo) Update(id uint64, newData model.Category) (category model.Category, err error) {
	var oldData model.Category
	tx := cr.db.First(&oldData, "id = ?", id)
	err = tx.Error
	if err != nil {
		return
	}

	// Modify data
	oldData.Update(newData)

	// Update data. Please keep in mind that this Save function
	// inserts data if no matching primary key found
	tx = cr.db.Save(&oldData)
	err = tx.Error

	return oldData, nil
}
