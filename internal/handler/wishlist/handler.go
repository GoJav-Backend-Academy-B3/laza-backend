package wishlist

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"

	usecase "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type getWishlistHandler struct {
	updateWishlistUsecase   usecase.UpdateWishListUsecase
	getWishlistUsecase      usecase.GetWishListUsecase
	getWishlistLimitUsecase usecase.GetWishListLimitUsecase
}

func (h *getWishlistHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs,
		handler.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/products/:id/wishlists",
			HandlerFunc: h.Put,
		},
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/wishlists",
			HandlerFunc: h.getByLimit,
		})
	return
}

func NewgetWishlistHandler(
	uws usecase.UpdateWishListUsecase,
	gws usecase.GetWishListUsecase,
	gls usecase.GetWishListLimitUsecase,

) handler.HandlerInterface {
	return &getWishlistHandler{
		updateWishlistUsecase:   uws,
		getWishlistUsecase:      gws,
		getWishlistLimitUsecase: gls,
	}
}
