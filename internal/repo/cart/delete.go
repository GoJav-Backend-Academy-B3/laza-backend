package cart

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *CartRepo) DeleteCart(model model.Cart) (value any, err error) {

	tx := r.db.Where("user_id = ? AND product_id = ? AND size_id = ?", model.UserId, model.ProductId, model.SizeId).
		Delete(&model)
	err = tx.Error
	value = "successfully delete product cart"
	return
}
