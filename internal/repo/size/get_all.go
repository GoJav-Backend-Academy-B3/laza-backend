package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) GetAll() (m []model.Size, err error) {
	tx := r.db.Find(&m)
	err = tx.Error
	return
}
