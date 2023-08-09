package brand

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type CreateBrandUsecase interface {
	Execute(request requests.BrandRequest) (brand model.Brand, err error)
}
