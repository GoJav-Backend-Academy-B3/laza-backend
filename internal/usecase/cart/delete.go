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

type deleteCartUsecase struct {
	deleteCartRepo d.DeleteAction[model.Cart]
	isCartRepo     dc.IsCarttByIdAction
}

func (us *deleteCartUsecase) Execute(userId, productId uint64) *helper.Response {

	tf := us.isCartRepo.IsCart(model.Cart{UserId: userId, ProductId: productId})
	if !tf {
		return helper.GetResponse(errors.New("there are no products in the cart").Error(), http.StatusNotFound, false)
	}

	id := map[string]uint64{"userId": userId, "productId": productId}
	err := us.deleteCartRepo.Delete(id)

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return helper.GetResponse("success", http.StatusOK, false)
}

func NewdeleteCartUsecase(deleteCartRepo d.DeleteAction[model.Cart], isCartRepo dc.IsCarttByIdAction) usecase.DeleteCartUsecase {
	return &deleteCartUsecase{
		deleteCartRepo: deleteCartRepo,
		isCartRepo:     isCartRepo,
	}
}
