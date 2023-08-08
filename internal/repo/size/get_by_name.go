package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) GetByName(size string) (m model.Size, err error) {
	tx := r.db.Where("size = ?", size).First(&m)
	err = tx.Error
	return
}
