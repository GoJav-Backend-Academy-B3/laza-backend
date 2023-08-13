package cart

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *CartRepo) UpdateCart(md model.Cart) (rs any, err error) {
	var _model model.Cart

	tx := r.db.Where("user_id=? AND product_id =? AND size_id = ?", md.UserId, md.ProductId, md.SizeId).
		Take(&_model)
	err = tx.Error

	if _model.Quantity == 1 {

		tx := r.db.
			Where("user_id=? AND product_id =? AND size_id =?", md.UserId, md.ProductId, md.SizeId).
			Delete(&_model)
		err = tx.Error
		rs = "successfully delete product cart"
	}

	if _model.Quantity > 1 {
		quantity := _model.Quantity - 1
		tx := r.db.
			Model(&_model).
			Where("user_id=? AND product_id =? AND size_id=?", md.UserId, md.ProductId, md.SizeId).
			Update("quantity", quantity)

		err = tx.Error
		rs = _model
	}

	return
}
