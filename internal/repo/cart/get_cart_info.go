package cart

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/response"
)

func (r *CartRepo) GetById(userId any) (rs response.CartOrderInfo, err error) {
	var subtotal float64
	r.db.Model(&model.Cart{}).Select("round(sum(cart.quantity * product.price),2) subtotal").Joins("left join product on cart.product_id = product.id").
		Where("cart.user_id = ?", userId).Scan(&subtotal)

	rs.SubTotal = subtotal
	rs.ShippingCost = 20000.00
	rs.Total = subtotal + rs.ShippingCost
	return
}
