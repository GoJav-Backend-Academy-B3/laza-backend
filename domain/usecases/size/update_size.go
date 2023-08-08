package size

import "github.com/phincon-backend/laza/domain/model"

type UpdateSizeUsecase interface {
	Execute(id uint64, m model.Size) (model.Size, error)
}
