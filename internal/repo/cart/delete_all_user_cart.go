package cart

import "github.com/phincon-backend/laza/domain/model"

func (r *CartRepo) Delete(userId any) (value string, err error) {
	user_id := userId.(uint64)
	tx := r.db.Where("user_id = ?", user_id).Delete(&model.Cart{})
	value = "successfully delete user cart"
	err = tx.Error
	return
}
