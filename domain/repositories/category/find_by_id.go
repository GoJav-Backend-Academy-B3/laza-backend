package category

import "github.com/phincon-backend/laza/domain/model"

type FindByIdAction interface {
	FindById(id uint64) (category model.Category, err error)
}
