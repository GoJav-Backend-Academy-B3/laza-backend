package cart

import (
	"errors"
	"net/http"

	dc "github.com/phincon-backend/laza/domain/repositories"
	d "github.com/phincon-backend/laza/domain/repositories/cart"
	"github.com/phincon-backend/laza/domain/response"
	r "github.com/phincon-backend/laza/domain/response"

	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
)

type getCartByIdUsecase struct {
	getCartByIdRepo   d.GetCartByIdAction
	cartOrderInfoRepo dc.GetByIdAction[r.CartOrderInfo]
}

func (uc *getCartByIdUsecase) Execute(userId uint64) *helper.Response {

	rs, err := uc.getCartByIdRepo.GetCartById(userId)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	if len(rs) == 0 {
		return helper.GetResponse(errors.New("there are no products in the cart").Error(), http.StatusNotFound, true)
	}

	cr, err := uc.cartOrderInfoRepo.GetById(userId)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	ci := response.CartInfo{CartPorduct: rs, CartOrderInfo: cr}

	return helper.GetResponse(ci, http.StatusOK, false)
}

func NewgetCartByIdUsecase(gcr d.GetCartByIdAction, coi dc.GetByIdAction[r.CartOrderInfo]) usecase.GetCartByIdUsecase {
	return &getCartByIdUsecase{
		getCartByIdRepo:   gcr,
		cartOrderInfoRepo: coi,
	}
}
