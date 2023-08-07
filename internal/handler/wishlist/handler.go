package wishlist

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"

	usecase "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type getWishlistHandler struct {
	updateWishlistUsecase usecase.UpdateWishListUsecase
	getWishlistUsecase    usecase.GetWishListUsecase
}

func (h *getWishlistHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs,
		handler.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/product/:productId/wishlist",
			HandlerFunc: h.Put,
		},
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/product/wishlist",
			HandlerFunc: h.get,
		})
	return
}

func NewgetWishlistHandler(
	uws usecase.UpdateWishListUsecase,
	gws usecase.GetWishListUsecase,
) handler.HandlerInterface {
	return &getWishlistHandler{
		updateWishlistUsecase: uws,
		getWishlistUsecase:    gws,
	}
}
