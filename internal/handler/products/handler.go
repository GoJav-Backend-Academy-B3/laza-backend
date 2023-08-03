package products

import (
	"net/http"

	handler "github.com/phincon-backend/laza/domain/handlers"
	usecase "github.com/phincon-backend/laza/domain/usecases/product"
)

type productHandler struct {
	path               string
	viewProductUsecase usecase.ViewProductUsecase
}

// GetHandlers implements handlers.HandlerInterface.
func (h *productHandler) GetHandlers() (hs []handler.HandlerStruct) {
	hs = append(hs, handler.HandlerStruct{
		Method:      http.MethodGet,
		Path:        h.path,
		HandlerFunc: h.get,
	})
	return
}

func NewProductHandler(
	path string,
	viewProductUsecase usecase.ViewProductUsecase) handler.HandlerInterface {
	return &productHandler{
		viewProductUsecase: viewProductUsecase,
	}
}
