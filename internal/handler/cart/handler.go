package cart

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/cart"
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
			Path:        "/products/:productId/cart",
			HandlerFunc: ch.post,
		},
		hd.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        "/products/:productId/cart",
			HandlerFunc: ch.Delete,
		},
		hd.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/products/:productId/cart",
			HandlerFunc: ch.Update,
		},
		hd.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/carts",
			HandlerFunc: ch.GetById,
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
