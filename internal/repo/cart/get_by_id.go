package cart

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/response"
)

func (r *CartRepo) GetCartById(userdId any) (rs []response.CartPorduct, err error) {

	err = r.db.Model(&model.Cart{}).
		Select("product.id, product.name , product.image_url, product.price, brand.name as brand_name, cart.quantity").
		Joins("left join product on cart.product_id = product.id left join brand on product.brand_id = brand.id").Where("cart.user_id =?", userdId).Scan(&rs).Error
	return
}
