package cart

import "github.com/phincon-backend/laza/domain/model"

func (r *CartRepo) Update(id any, cr model.Cart) (rs model.Cart, err error) {

	if r.db.Where("user_id=? AND product_id =?", cr.UserId, cr.ProductId).Find(&cr); rs.Quantity == 1 {
		err = r.db.Where("user_id=? AND product_id =?", cr.UserId, cr.ProductId).Delete(&cr).Scan(&rs).Error
	} else {
		if cr.Quantity > 1 {
			quantity := cr.Quantity - 1
			err = r.db.Model(&cr).Where("user_id=? AND product_id =?", cr.UserId, cr.ProductId).Update("quantity", quantity).Scan(&rs).Error
		}
	}
	return
}
