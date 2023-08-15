package cart

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	dc "github.com/phincon-backend/laza/domain/repositories/cart"
	"github.com/phincon-backend/laza/domain/requests"

	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
)

type deleteCartUsecase struct {
	deleteCartRepo dc.DeleteCartAction
	validate       *validator.Validate
}

func (us *deleteCartUsecase) Execute(userId uint64, rb requests.CartRequest) (value any, err error) {
	err = us.validate.Struct(rb)

	if err != nil {
		return
	}

	_model := model.Cart{UserId: userId, ProductId: rb.ProductId, SizeId: rb.SizeId}
	value, err = us.deleteCartRepo.DeleteCart(_model)

	return value, err
}

func NewdeleteCartUsecase(
	deleteCartRepo dc.DeleteCartAction,
	validator *validator.Validate,
) usecase.DeleteCartUsecase {
	return &deleteCartUsecase{
		deleteCartRepo: deleteCartRepo,
		validate:       validator,
	}
}
