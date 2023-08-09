package response

import "github.com/phincon-backend/laza/domain/model"

type CategorySimpleResponse struct {
	Id       uint64 `json:"id"`
	Category string `json:"category"`
}

func (p CategorySimpleResponse) FillFromEntity(e model.Category) CategorySimpleResponse {
	p.Id = e.Id
	p.Category = e.Category
	return p
}
