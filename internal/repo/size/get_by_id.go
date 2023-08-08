package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) GetById(id_r any) (size model.Size, err error) {

	id := id_r.(uint64)
	// Grab size that matches with `id`
	tx := r.db.First(&size, "id = ?", id)
	// TODO: result.Error might return gorm.ErrRecordNotFound if
	//       no `id` can be found
	err = tx.Error

	return size, err
}
