package cart

import (
	"errors"
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"

	h "github.com/phincon-backend/laza/helper"
)

type insertCartUsecase struct {
	insertCartRepo       d.InsertAction[m.Cart]
	getProductByIdAction d.GetByIdAction[m.Product]
}

func (uc *insertCartUsecase) Execute(userId uint64, productId uint64) *h.Response {

	_, err := uc.getProductByIdAction.GetById(productId)

	if err == gorm.ErrRecordNotFound {
		return helper.GetResponse(errors.New("product not found").Error(), http.StatusNotFound, true)
	}

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	cart := m.Cart{
		UserId:    userId,
		ProductId: productId,
		Quantity:  1,
	}
	rs, err := uc.insertCartRepo.Insert(cart)
	if err != nil {
		return h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return h.GetResponse(rs, http.StatusOK, false)
}

func NewinsertCartUsecase(
	icp d.InsertAction[m.Cart],
	gpi d.GetByIdAction[m.Product],
) usecase.InsertCartUsecase {
	return &insertCartUsecase{
		insertCartRepo:       icp,
		getProductByIdAction: gpi,
	}
}
