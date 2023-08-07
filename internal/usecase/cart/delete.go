package cart

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
)

type deleteCartUsecase struct {
	deleteCartRepo d.DeleteAction[model.Cart]
}

func (us *deleteCartUsecase) Execute(userId, productId uint64) *helper.Response {
	id := map[string]uint64{"userId": userId, "productId": productId}
	err := us.deleteCartRepo.Delete(id)

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return helper.GetResponse("success", http.StatusOK, false)
}

func NewdeleteCartUsecase(deleteCartRepo d.DeleteAction[model.Cart]) usecase.DeleteCartUsecase {
	return &deleteCartUsecase{
		deleteCartRepo: deleteCartRepo,
	}
}
