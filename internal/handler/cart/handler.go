package cart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/cart"
	"github.com/phincon-backend/laza/middleware"
)

type CartHandler struct {
	insertCartUc   uc.InsertCartUsecase
	deleteCartUc   uc.DeleteCartUsecase
	updateCartUc   uc.UpdateCartUsecase
	getCartByIdUc  uc.GetCartByIdUsecase
	getCartOrderUc uc.GetCartOrderInfoUsecase
}

func (ch *CartHandler) GetHandlers() (h []hd.HandlerStruct) {
	h = append(h,
		hd.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/carts",
			HandlerFunc: ch.post,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
		},
		hd.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        "/carts",
			HandlerFunc: ch.Delete,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
		},
		hd.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/carts",
			HandlerFunc: ch.Update,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
		},
		hd.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/carts",
			HandlerFunc: ch.GetById,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware()},
		},
	)
	return h
}

func NewcartHandler(
	insertCartUc uc.InsertCartUsecase,
	deleteCartUc uc.DeleteCartUsecase,
	updateCartUc uc.UpdateCartUsecase,
	getCartByIdUc uc.GetCartByIdUsecase,

) hd.HandlerInterface {
	return &CartHandler{
		insertCartUc:  insertCartUc,
		deleteCartUc:  deleteCartUc,
		updateCartUc:  updateCartUc,
		getCartByIdUc: getCartByIdUc,
	}
}
