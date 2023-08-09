package brand

import "github.com/phincon-backend/laza/domain/model"

type GetBrandByIdUsecase interface {
	Execute(brandId uint64) (brand model.Brand, err error)
}
