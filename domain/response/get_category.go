package response

import "github.com/phincon-backend/laza/domain/model"

type GetCategory struct {
	Id       uint64 `json:"id"`
	Category string `json:"category"`
}

func (p GetCategory) FillFromEntity(e model.Category) GetCategory {
	p.Id = e.Id
	p.Category = e.Category
	return p
}
