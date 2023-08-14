package brand

import (
	"context"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *BrandRepo) GetByName(brand string) (m model.Brand, err error) {
	db := r.db.Where("name = ?", brand).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}

func (r *BrandRepo) GetByNameWithContext(ctx context.Context, brand string) (m model.Brand, err error) {
	db := r.db.WithContext(ctx).Where("name = ?", brand).First(&m)
	err = db.Error
	// TODO: Should return error here
	return
}
