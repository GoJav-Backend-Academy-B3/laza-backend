package response

import "github.com/phincon-backend/laza/domain/model"

type Size struct {
	Id   uint64 `json:"id"`
	Size string `json:"size"`
}
type GetSize struct {
	Id   uint64 `json:"id"`
	Size string `json:"size"`
}

func (p GetSize) FillFromEntity(e model.Size) GetSize {
	p.Id = e.Id
	p.Size = e.Size
	return p
}

func (s *Size) FillFromEntity(m model.Size) {
	s.Id = m.Id
	s.Size = m.Size
}
