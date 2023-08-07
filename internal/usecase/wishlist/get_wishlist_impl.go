package wishlist

import (
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	r "github.com/phincon-backend/laza/domain/response"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
	h "github.com/phincon-backend/laza/helper"
)

type getWishlistUsecase struct {
	getWishlistRepo   d.GetByIdAction[[]m.Wishlist]
	getAllProductRepo d.GetAllAction[m.Product]
}

func NewgetWishlistUsecase(gwp d.GetByIdAction[[]m.Wishlist],
	gap d.GetAllAction[m.Product]) p.GetWishListUsecase {
	return &getWishlistUsecase{
		getWishlistRepo:   gwp,
		getAllProductRepo: gap,
	}
}

func (uc *getWishlistUsecase) Execute(userId uint64) *h.Response {
	ws, err := uc.getWishlistRepo.GetById(userId)
	if err != nil {
		h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	pd, err := uc.getAllProductRepo.GetAll()
	if err != nil {
		h.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	wp := r.WishlistProduct{}
	for _, ws := range ws {
		for _, pd := range pd {
			if ws.ProductId == pd.Id {
				rp := r.Product{}
				rp.FillFromEntity(pd)
				wp.Products = append(wp.Products, rp)
			}
		}
	}
	wp.Total = len(wp.Products)
	return h.GetResponse(wp, http.StatusOK, false)
}
