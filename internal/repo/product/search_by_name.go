package product

import (
	"fmt"

	"github.com/phincon-backend/laza/domain/model"
)

func (p *ProductRepo) SearchByName(keyword string, offset, limit uint64) (ms []model.Product, err error) {
	tx := p.db.Where("lower(name) LIKE lower(?)", fmt.Sprintf("%%%s%%", keyword)).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&ms)
	err = tx.Error
	return
}
