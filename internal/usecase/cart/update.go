package cart

import (
	"errors"
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	dc "github.com/phincon-backend/laza/domain/repositories/cart"

	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
)

type updateCartUsecase struct {
	updateCartRepo d.UpdateAction[model.Cart]
	isCartRepo     dc.IsCarttByIdAction
}

func (uc *updateCartUsecase) Execute(userId, productId uint64) *helper.Response {
	md := model.Cart{UserId: userId, ProductId: productId}

	if bl := uc.isCartRepo.IsCart(md); !bl {
		return helper.GetResponse(errors.New("There are no products in the cart").Error(), http.StatusNotFound, true)
	}

	rs, err := uc.updateCartRepo.Update("", md)

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return helper.GetResponse(rs, http.StatusOK, false)
}

func NewupdateCartUsecase(updateCartRepo d.UpdateAction[model.Cart], isCartRepo dc.IsCarttByIdAction) usecase.UpdateCartUsecase {
	return &updateCartUsecase{
		updateCartRepo: updateCartRepo,
		isCartRepo:     isCartRepo,
	}
}
