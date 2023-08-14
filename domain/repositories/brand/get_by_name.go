package brand

import (
	"context"

	"github.com/phincon-backend/laza/domain/model"
)

type GetByNameAction interface {
	GetByName(brand string) (model.Brand, error)
}

type GetByNameActionWithContext interface {
	GetByNameWithContext(ctx context.Context, brand string) (model.Brand, error)
}
