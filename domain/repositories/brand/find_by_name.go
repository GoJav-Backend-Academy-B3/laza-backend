package brand

import "github.com/phincon-backend/laza/domain/model"

type FindByNameAction interface {
	FindByName(name string) (brands []model.Brand, err error)
}
