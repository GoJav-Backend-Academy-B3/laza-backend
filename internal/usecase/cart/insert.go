package cart

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
)

type insertCartUsecase struct {
	insertCartRepo d.InsertAction[m.Cart]
	validate       *validator.Validate
}

func (uc *insertCartUsecase) Execute(userid uint64, rb requests.CartRequest) (_result model.Cart, err error) {

	err = uc.validate.Struct(rb)
	if err != nil {
		return
	}
	_model := model.Cart{
		UserId:    userid,
		ProductId: rb.ProductId,
		SizeId:    rb.SizeId,
		Quantity:  1,
	}

	_result, err = uc.insertCartRepo.Insert(_model)
	return

}

func NewinsertCartUsecase(
	icp d.InsertAction[m.Cart],
	validate *validator.Validate,

) usecase.InsertCartUsecase {
	return &insertCartUsecase{
		insertCartRepo: icp,
		validate:       validate,
	}
}
