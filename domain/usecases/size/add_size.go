package size

import "github.com/phincon-backend/laza/domain/model"

type AddSizeUsecase interface {
	Execute(name string) (model.Size, error)
}
