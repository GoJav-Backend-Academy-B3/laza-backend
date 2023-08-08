package category

import "github.com/phincon-backend/laza/domain/model"

type GetByNameAction interface {
	GetByName(category string) (model.Category, error)
}
