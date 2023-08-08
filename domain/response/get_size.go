package response

import "github.com/phincon-backend/laza/domain/model"

type Size struct {
	Id   uint64 `json:"id"`
	Size string `json:"size"`
}

func (s *Size) FillFromEntity(m model.Size) {
	s.Id = m.Id
	s.Size = m.Size
}
