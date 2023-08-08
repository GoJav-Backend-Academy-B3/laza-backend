package category

import "github.com/phincon-backend/laza/domain/model"

type FindByNameAction interface {
	FindByName(name string) (categories []model.Category, err error)
}
