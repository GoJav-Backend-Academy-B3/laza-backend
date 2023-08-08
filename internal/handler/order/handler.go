package order

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/order"
	"net/http"
)

type orderHandler struct {
	createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase
	createOrderWithBankUsecase  order.CreateOrderWithBankUsecase
}

func NewOrderHandler(createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase, createOrderWithBankUsecase order.CreateOrderWithBankUsecase) *orderHandler {
	return &orderHandler{createOrderWithGopayUsecase: createOrderWithGopayUsecase, createOrderWithBankUsecase: createOrderWithBankUsecase}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *orderHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/order/gopay",
			HandlerFunc: h.CreateOrderWithGopay,
		}, handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/order/bank",
			HandlerFunc: h.CreateOrderWithBank,
		},
	)
	return
}
