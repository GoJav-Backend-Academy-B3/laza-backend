package user

import "github.com/phincon-backend/laza/domain/response"

func (r *UserRepo) GetAll() (es []response.User, err error) {
	tx := r.db.Find(&es)
	err = tx.Error
	return
}
