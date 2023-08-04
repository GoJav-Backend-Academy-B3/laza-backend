package wishlist

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"

	usecase "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type getWishlistHandler struct {
	path                  string
	updateWishlistUsecase usecase.UpdateWishListUsecase
}

func (h *getWishlistHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs, handler.HandlerStruct{
		Method:      http.MethodPut,
		Path:        h.path,
		HandlerFunc: h.Put,
	})
	return
}

func NewgetWishlistHandler(path string,
	uws usecase.UpdateWishListUsecase) handler.HandlerInterface {
	return &getWishlistHandler{
		path:                  path,
		updateWishlistUsecase: uws,
	}
}
