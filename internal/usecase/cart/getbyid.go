package cart

import (
	"errors"
	"net/http"

	d "github.com/phincon-backend/laza/domain/repositories/cart"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
)

type getCartByIdUsecase struct {
	getCartByIdRepo d.GetCartByIdAction
}

func (uc *getCartByIdUsecase) Execute(userId any) *helper.Response {

	rs, err := uc.getCartByIdRepo.GetCartById(userId)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	if len(rs) == 0 {
		return helper.GetResponse(errors.New("there are no products in the cart").Error(), http.StatusNotFound, true)
	}

	return helper.GetResponse(rs, http.StatusOK, false)
}

func NewgetCartByIdUsecase(gcr d.GetCartByIdAction) usecase.GetCartByIdUsecase {
	return &getCartByIdUsecase{
		getCartByIdRepo: gcr,
	}
}
