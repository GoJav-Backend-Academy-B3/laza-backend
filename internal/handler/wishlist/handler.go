package wishlist

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handler "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/middleware"

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
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
		},
		handler.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/wishlists",
			HandlerFunc: h.getByLimit,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
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
