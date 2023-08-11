package category

import (
	"context"

	"github.com/phincon-backend/laza/domain/model"
)

type GetByNameAction interface {
	GetByName(category string) (model.Category, error)
}

type GetByNameActionWithContext interface {
	GetByNameWithContext(ctx context.Context, category string) (model.Category, error)
}
