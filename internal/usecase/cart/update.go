package cart

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	dc "github.com/phincon-backend/laza/domain/repositories/cart"
	"github.com/phincon-backend/laza/domain/requests"

	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
)

type updateCartUsecase struct {
	updateCartRepo dc.UpdateCartAction
	validate       *validator.Validate
}

func (uc *updateCartUsecase) Execute(userId uint64, rb requests.CartRequest) (_result any, err error) {
	err = uc.validate.Struct(rb)
	if err != nil {
		return
	}

	_model := model.Cart{UserId: userId, ProductId: rb.ProductId, SizeId: rb.SizeId}
	_result, err = uc.updateCartRepo.UpdateCart(_model)

	return

}

func NewupdateCartUsecase(
	updateCartRepo dc.UpdateCartAction,
	validate *validator.Validate,
) usecase.UpdateCartUsecase {
	return &updateCartUsecase{
		updateCartRepo: updateCartRepo,
		validate:       validate,
	}
}
