package cart

import (
	"github.com/phincon-backend/laza/helper"
)

type InsertCartUsecase interface {
	Execute(userid uint64, productId uint64) *helper.Response
}
