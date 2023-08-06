package cart

import (
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
	h "github.com/phincon-backend/laza/helper"
)

type insertCartUsecase struct {
	insertCartRepo d.InsertAction[m.Cart]
}

func (uc *insertCartUsecase) Execute(userId uint64, productId uint64) *h.Response {
	cart := m.Cart{
		UserId:    userId,
		ProductId: productId,
		Quantity:  1,
	}

	// check the product first
	rs, err := uc.insertCartRepo.Insert(cart)
	if err != nil {
		return h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	return h.GetResponse(rs, http.StatusOK, false)
}

func NewinsertCartUsecase(icp d.InsertAction[m.Cart]) usecase.InsertCartUsecase {
	return &insertCartUsecase{
		insertCartRepo: icp,
	}
}
