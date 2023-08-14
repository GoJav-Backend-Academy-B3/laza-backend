package category

import (
	"context"

	"github.com/phincon-backend/laza/domain/model"
)

func (cr *CategoryRepo) GetByName(category string) (m model.Category, err error) {
	db := cr.db.Where("category = ?", category).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}

func (cr *CategoryRepo) GetByNameWithContext(ctx context.Context, category string) (m model.Category, err error) {
	db := cr.db.WithContext(ctx).Where("category = ?", category).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}
