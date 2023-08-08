package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) Delete(id_r any) error {

	id := id_r.(uint64)

	var size model.Size
	tx := r.db.First(&size, "id = ?", id)
	err := tx.Error
	if err != nil {
		return err
	}

	tx = r.db.Delete(&size)
	err = tx.Error
	return err
}
