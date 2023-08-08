package size

import "github.com/phincon-backend/laza/domain/model"

type GetAllSizeUsecase interface {
	Execute() ([]model.Size, error)
}
