package cart

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/response"
)

func (r *CartRepo) GetCartOrderInfo(userId any, rs *response.CartOrderInfo) {
	var subtotal float64
	r.db.Model(&model.Cart{}).Select("(cart.quantity * product.price) subtotal").Joins("left join product on cart.product_id = product.id").
		Where("cart.user_id = ?", userId).Scan(&subtotal)

	rs.SubTotal = subtotal
	rs.Total = subtotal
	return
}
