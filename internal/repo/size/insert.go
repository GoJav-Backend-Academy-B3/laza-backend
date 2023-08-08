package size

import "github.com/phincon-backend/laza/domain/model"

func (r *SizeRepo) Insert(e model.Size) (model.Size, error) {
	tx := r.db.Create(&e)
	return e, tx.Error
}
