package cart

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
)

type updateCartUsecase struct {
	updateCartRepo d.UpdateAction[model.Cart]
}

func (uc *updateCartUsecase) Execute(userId, productId uint64) *helper.Response {
	rs, err := uc.updateCartRepo.Update("", model.Cart{UserId: userId, ProductId: productId})

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	if rs == (model.Cart{}) {
		return helper.GetResponse("there are no products in the cart", http.StatusNotFound, false)
	}

	return helper.GetResponse(rs, http.StatusOK, false)
}

func NewupdateCartUsecase(updateCartRepo d.UpdateAction[model.Cart]) usecase.UpdateCartUsecase {
	return &updateCartUsecase{
		updateCartRepo: updateCartRepo,
	}
}
