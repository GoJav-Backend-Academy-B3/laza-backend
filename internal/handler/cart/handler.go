package cart

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/cart"
)

type CartHandler struct {
	insertCartUc uc.InsertCartUsecase
	deleteCartUc uc.DeleteCartUsecase
}

func (ch *CartHandler) GetHandlers() (h []hd.HandlerStruct) {
	h = append(h,
		hd.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/product/:productId/cart",
			HandlerFunc: ch.post,
		},
		hd.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        "/product/:productId/cart",
			HandlerFunc: ch.Delete,
		},
	)
	return h
}

func NewcartHandler(
	insertCartUc uc.InsertCartUsecase,
	deleteCartUc uc.DeleteCartUsecase,

) hd.HandlerInterface {
	return &CartHandler{
		insertCartUc: insertCartUc,
		deleteCartUc: deleteCartUc,
	}
}
