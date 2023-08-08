package category

import (
	"fmt"
	"github.com/phincon-backend/laza/domain/model"
)

func (cr *CategoryRepo) FindByName(name string) (categories []model.Category, err error) {
	db := cr.db.Where("lower(category) LIKE lower(?)", fmt.Sprintf("%%%s%%", name)).Find(&categories)
	err = db.Error
	return
}
