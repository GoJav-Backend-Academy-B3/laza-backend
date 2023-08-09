package brand

import "github.com/phincon-backend/laza/domain/model"

type ViewBrandUsecase interface {
	Execute() (brands []model.Brand, err error)
}
