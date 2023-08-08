package category

import "github.com/phincon-backend/laza/domain/model"

type FindAllAction interface {
	FindAll() (categories []model.Category, err error)
}
