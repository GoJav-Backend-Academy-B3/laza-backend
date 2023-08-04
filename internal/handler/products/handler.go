package products

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
)

type productHandler struct {
	path               string
	viewProductUsecase         uc.ViewProductUsecase
}

// GetHandlers implements handlers.HandlerInterface.
	hs = append(hs, handler.HandlerStruct{
func (h *productHandler) GetHandlers() (hs []hd.HandlerStruct) {
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodGet,
		Path:        h.path,
		HandlerFunc: h.get,
	})
	return
}

func NewProductHandler(
	path string,
	viewProductUsecase usecase.ViewProductUsecase) handler.HandlerInterface {
	viewProductUsecase uc.ViewProductUsecase,
	return &productHandler{
		path:               path,
		viewProductUsecase: viewProductUsecase,
	}
}
