package brand

import "github.com/phincon-backend/laza/domain/model"

type SearchBrandByNameUsecase interface {
	Execute(brandName string) (brands []model.Brand, err error)
}
