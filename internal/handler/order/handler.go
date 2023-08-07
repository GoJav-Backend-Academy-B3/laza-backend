package order

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/order"
	"net/http"
)

type orderHandler struct {
	createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase
	validate                    *validator.Validate
}

func NewOrderHandler(createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase, validate *validator.Validate) *orderHandler {
	return &orderHandler{createOrderWithGopayUsecase: createOrderWithGopayUsecase, validate: validate}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *orderHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs, handlers.HandlerStruct{
		Method:      http.MethodPost,
		Path:        "/order/gopay",
		HandlerFunc: h.CreateOrderWithGopay,
	},
	)
	return
}
