package products

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
)

type productHandler struct {
	path                       string
	viewProductUsecase         uc.ViewProductUsecase
	searchProductByNameUsecase uc.SearchProductByNameUsecase
}

// GetHandlers implements handlers.HandlerInterface.
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
	viewProductUsecase uc.ViewProductUsecase,
	searchProductByNameUsecase uc.SearchProductByNameUsecase) hd.HandlerInterface {
	return &productHandler{
		path:                       path,
		viewProductUsecase:         viewProductUsecase,
		searchProductByNameUsecase: searchProductByNameUsecase,
	}
}
