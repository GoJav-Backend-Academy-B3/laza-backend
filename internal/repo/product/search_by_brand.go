package product

import (
	"github.com/phincon-backend/laza/domain/model"
	"strings"
)

func (p *ProductRepo) SearchByBrand(brand string, offset, limit uint64) (ms []model.Product, err error) {
	tx := p.db.Table("product p").
		Select("p.name, p.price, p.image_url, p.description, p.brand_id, p.created_at, p.updated_at, p.id").
		Joins("JOIN brand ON p.brand_id = brand.id").
		Where("lower(brand.name) = ?", strings.ToLower(brand)).
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&ms)

	err = tx.Error
	return
}
