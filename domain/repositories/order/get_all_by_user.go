package order

import "github.com/phincon-backend/laza/domain/model"

type GetAllByUser interface {
	GetAllByUser(userId uint64) (order []model.Order, err error)
}
