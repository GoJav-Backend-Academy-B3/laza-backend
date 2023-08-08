package size

import "github.com/phincon-backend/laza/domain/model"

type GetSizeById interface {
	Execute(id uint64) (model.Size, error)
}
