package product

import "github.com/phincon-backend/laza/domain/entities"

func (r *ProductRepo) GetWithLimit(limit, offset uint64) (es []entities.Product, err error) {
	tx := r.db.Find(&es).Offset(int(offset)).Limit(int(limit))
	err = tx.Error
	return
}
