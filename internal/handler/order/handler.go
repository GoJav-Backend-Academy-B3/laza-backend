package order

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/order"
	"net/http"
)

type orderHandler struct {
	createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase
	createOrderWithBankUsecase  order.CreateOrderWithBankUsecase
	createOrderWithCCUsecase    order.CreateOrderWithCCUsecase
	getById                     order.GetOrderByIdUsecase
}

func NewOrderHandler(
	createOrderWithGopayUsecase order.CreateOrderWithGopayUsecase,
	createOrderWithBankUsecase order.CreateOrderWithBankUsecase,
	createOrderWithCCUsecase order.CreateOrderWithCCUsecase,
	getById order.GetOrderByIdUsecase,
) *orderHandler {
	return &orderHandler{
		createOrderWithGopayUsecase: createOrderWithGopayUsecase,
		createOrderWithBankUsecase:  createOrderWithBankUsecase,
		createOrderWithCCUsecase:    createOrderWithCCUsecase,
		getById:                     getById,
	}
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
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/order/cc",
			HandlerFunc: h.CreateOrderWithCC,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/order/:order_id",
			HandlerFunc: h.GetOrderById,
		},
	)
	return
}
