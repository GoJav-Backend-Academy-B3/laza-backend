package brand

import (
	"fmt"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *BrandRepo) FindByName(name string) (brands []model.Brand, err error) {
	db := r.db.Where("lower(category) LIKE lower(?)", fmt.Sprintf("%%%s%%", name)).Find(&brands)
	err = db.Error
	return
}
