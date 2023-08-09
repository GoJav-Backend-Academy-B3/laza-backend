package brand

import "github.com/phincon-backend/laza/domain/model"

type GetByNameAction interface {
	GetByName(brand string) (model.Brand, error)
}
