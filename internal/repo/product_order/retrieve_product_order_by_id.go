package product_order

import "github.com/phincon-backend/laza/domain/model"

func (r *ProductOrderRepo) GetById(orderId any) (po model.ProductOrder, err error) {
	orderIdConv := orderId.(string)

	tx := r.db.First(&po, orderIdConv)
	err = tx.Error
	return
}
