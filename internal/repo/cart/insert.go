package cart

import "github.com/phincon-backend/laza/domain/model"

func (r *CartRepo) Insert(cart model.Cart) (cr model.Cart, err error) {

	if r.db.
		Where("user_id =? AND product_id =? AND size_id = ?", cart.UserId, cart.ProductId, cart.SizeId).
		Take(&cr); cr == (model.Cart{}) {
		tx := r.db.Create(&cart)
		err = tx.Error
		cr = cart
	} else {
		tx := r.db.Model(&cr).
			Where("user_id =? AND product_id =? AND size_id = ?", cart.UserId, cart.ProductId, cart.SizeId).
			Update("quantity", cr.Quantity+1)
		err = tx.Error
	}
	return
}
