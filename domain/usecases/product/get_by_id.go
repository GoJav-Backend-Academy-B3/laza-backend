package product

import "github.com/phincon-backend/laza/helper"

type GetByIdProductUsecase interface {
	Execute(id uint64) *helper.Response
}
