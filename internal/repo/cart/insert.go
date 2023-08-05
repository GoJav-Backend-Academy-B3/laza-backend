package cart

import "github.com/phincon-backend/laza/domain/model"

func (r *CartRepo) Insert(cart model.Cart) (cr model.Cart, err error) {

	if r.db.Where("user_id =? AND product_id =?", cart.UserId, cart.ProductId).Find(&cr); cr == (model.Cart{}) {
		err = r.db.Create(&cart).Scan(&cr).Error
		if err != nil {
			return
		}
	} else {
		err = r.db.Model(&cr).Where("user_id =? AND product_id =?", cart.UserId, cart.ProductId).
			Update("quantity", cr.Quantity+1).Error
		if err != nil {
			return
		}
	}
	return
}
