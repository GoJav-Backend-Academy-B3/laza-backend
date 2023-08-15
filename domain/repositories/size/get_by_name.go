package size

import (
	"github.com/phincon-backend/laza/domain/model"
	"golang.org/x/net/context"
)

type GetByNameAction interface {
	GetByName(size string) (model.Size, error)
}

type GetByNameActionWithContext interface {
	GetByNameWithContext(ctx context.Context, size string) (model.Size, error)
}
