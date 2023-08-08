package wishlist

import (
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	dr "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/response"

	d "github.com/phincon-backend/laza/domain/repositories/wishlist"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
	h "github.com/phincon-backend/laza/helper"
)

type getWishlistLimitUsecase struct {
	getWishlistLimitRepo d.GetCartWithLimitAction
	getWishlistRepo      dr.GetByIdAction[[]m.Wishlist]
}

func (uc *getWishlistLimitUsecase) Execute(userId, offset, limit uint64) *h.Response {

	rs, err := uc.getWishlistLimitRepo.GetCartWithLimit(userId, offset, limit)
	if err != nil {
		h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	cr, err := uc.getWishlistRepo.GetById(userId)
	if err != nil {
		h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	md := response.WishListProductLimit{
		Total:           len(cr),
		WishlistProduct: rs,
	}

	return h.GetResponse(md, http.StatusOK, false)

}

func NewgetWishlistLimitUsecase(gwr d.GetCartWithLimitAction, cr dr.GetByIdAction[[]m.Wishlist]) p.GetWishListLimitUsecase {
	return &getWishlistLimitUsecase{
		getWishlistLimitRepo: gwr,
		getWishlistRepo:      cr,
	}
}
