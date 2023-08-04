package product

import (
	"fmt"

	"github.com/phincon-backend/laza/domain/model"
)

func (p *ProductRepo) SearchByName(keyword string, offset, limit uint64) (ms []model.Product, err error) {
	tx := p.db.Where("lower(name) LIKE lower(?)", fmt.Sprintf("%%%s%%", keyword)).
		Find(&ms).
		Offset(int(offset)).
		Limit(int(limit))
	err = tx.Error
	return
}
