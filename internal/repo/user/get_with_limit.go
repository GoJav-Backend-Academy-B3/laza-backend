package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) GetWithLimit(offset, limit uint64) (es []model.User, err error) {
	tx := r.db.Offset(int(offset)).Limit(int(limit)).Find(&es)
	err = tx.Error
	return
}
