package cart

import (
	"net/http"

	d "github.com/phincon-backend/laza/domain/repositories"
	m "github.com/phincon-backend/laza/domain/response"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	h "github.com/phincon-backend/laza/helper"
)

type getCartOrderInfoUsecase struct {
	cartOrderInfoRepo d.GetByIdAction[m.CartOrderInfo]
}

func (co *getCartOrderInfoUsecase) Execute(userId uint64) *h.Response {
	rs, err := co.cartOrderInfoRepo.GetById(userId)

	if err != nil {
		h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return h.GetResponse(rs, http.StatusOK, false)
}

func NewgetCartOrderInfoUsecase(cartOrderInfoRepo d.GetByIdAction[m.CartOrderInfo]) usecase.GetCartOrderInfoUsecase {
	return &getCartOrderInfoUsecase{
		cartOrderInfoRepo: cartOrderInfoRepo,
	}
}
