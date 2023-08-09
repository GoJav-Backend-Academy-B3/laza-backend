package response

import "github.com/phincon-backend/laza/domain/model"

type Size struct {
	Id   uint64 `json:"id"`
	Size string `json:"size"`
}

func (p *Size) FillFromEntity(e model.Size) {
	p.Id = e.Id
	p.Size = e.Size

}
