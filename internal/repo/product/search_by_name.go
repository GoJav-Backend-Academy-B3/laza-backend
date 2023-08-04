package product

import "github.com/phincon-backend/laza/domain/model"

func (p *ProductRepo) SearchByName(keyword string, offset, limit uint64) (ms []model.Product, err error) {
	tx := p.db.Where("name LIKE CONCAT('%', ?, '%')", keyword).
		Find(&ms).
		Offset(int(offset)).
		Limit(int(limit))
	err = tx.Error
	return
}
