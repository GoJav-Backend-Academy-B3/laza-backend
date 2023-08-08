package category

import "github.com/phincon-backend/laza/domain/model"

type UpdateByIdAction interface {
	Update(id uint64, model model.Category) (category model.Category, err error)
}
