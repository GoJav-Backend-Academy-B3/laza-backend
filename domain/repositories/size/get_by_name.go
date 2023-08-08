package size

import "github.com/phincon-backend/laza/domain/model"

type GetByNameAction interface {
	GetByName(size string) (model.Size, error)
}
