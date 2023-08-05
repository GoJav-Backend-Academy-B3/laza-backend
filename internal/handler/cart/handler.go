package cart

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/cart"
)

type CartHandler struct {
	insertCartUc uc.InsertCartUsecase
}

func (ch *CartHandler) GetHandlers() (h []hd.HandlerStruct) {
	h = append(h,
		hd.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/product/:productId/cart",
			HandlerFunc: ch.post,
		},
	)
	return h
}

func NewcartHandler(
	insertCartUc uc.InsertCartUsecase,
) hd.HandlerInterface {
	return &CartHandler{
		insertCartUc: insertCartUc,
	}
}
