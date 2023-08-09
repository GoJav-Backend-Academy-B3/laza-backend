package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type UpdateBrandNameByIdUsecase interface {
	Execute(id uint64, request requests.BrandRequest) (brand model.Brand, err error)
}
