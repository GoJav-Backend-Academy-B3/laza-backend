package user

import (
	"github.com/phincon-backend/laza/domain/response"
)

func (r *UserRepo) GetWithLimit(offset, limit uint64) (es []response.User, err error) {
	tx := r.db.Offset(int(offset)).Limit(int(limit)).Find(&es)
	err = tx.Error
	return
}
