package response

import "github.com/phincon-backend/laza/domain/model"

type GetSize struct {
	Id   uint64 `json:"id"`
	Size string `json:"size"`
}

func (p GetSize) FillFromEntity(e model.Size) GetSize {
	p.Id = e.Id
	p.Size = e.Size
	return p
}
