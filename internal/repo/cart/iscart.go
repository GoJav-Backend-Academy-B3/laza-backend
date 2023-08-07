package cart

import "github.com/phincon-backend/laza/domain/model"

func (r *CartRepo) IsCart(cr model.Cart) (rs bool) {
	r.db.Find(&cr)
	if cr.Quantity == 0 {
		return false
	}
	return true
}
