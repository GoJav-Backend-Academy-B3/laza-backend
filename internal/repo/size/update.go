package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) Update(id_r any, e model.Size) (size model.Size, err error) {

	id := id_r.(uint64)
	// Grab size that matches with `id`
	tx := r.db.First(&size, "id = ?", id)
	// TODO: result.Error might return gorm.ErrRecordNotFound if
	//       no `id` can be found
	err = tx.Error
	if err != nil {
		return
	}

	size.Update(e)

	// Update data. Please keep in mind that this Save function
	// inserts data if no matching primary key found
	tx = r.db.Save(&size)
	err = tx.Error

	return size, err
}
