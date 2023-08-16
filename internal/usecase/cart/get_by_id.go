package cart

import (
	"fmt"

	dc "github.com/phincon-backend/laza/domain/repositories"
	d "github.com/phincon-backend/laza/domain/repositories/cart"
	"github.com/phincon-backend/laza/domain/response"
	r "github.com/phincon-backend/laza/domain/response"

	usecase "github.com/phincon-backend/laza/domain/usecases/cart"
)

type getCartByIdUsecase struct {
	getCartByIdRepo   d.GetCartByIdAction
	cartOrderInfoRepo dc.GetByIdAction[r.CartOrderInfo]
}

func (uc *getCartByIdUsecase) Execute(userId uint64) (_result response.CartInfo, err error) {

	rs, err := uc.getCartByIdRepo.GetCartById(userId)
	if err != nil {
		return
	}
	cr, err := uc.cartOrderInfoRepo.GetById(userId)
	if err != nil {
		return
	}

	// set cost value
	if len(rs) == 0 {
		cr.ShippingCost = 0
		cr.Total = 0
		fmt.Println("GGWP")
	}
	_result = response.CartInfo{CartPorduct: rs, CartOrderInfo: cr}

	return
}

func NewgetCartByIdUsecase(gcr d.GetCartByIdAction, coi dc.GetByIdAction[r.CartOrderInfo]) usecase.GetCartByIdUsecase {
	return &getCartByIdUsecase{
		getCartByIdRepo:   gcr,
		cartOrderInfoRepo: coi,
	}
}
