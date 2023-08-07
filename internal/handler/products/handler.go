package products

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
)

type productHandler struct {
	path                       string
	viewProductUsecase         uc.ViewProductUsecase
	searchProductByNameUsecase uc.SearchProductByNameUsecase
	getByIdProduct             uc.GetByIdProductUsecase
}

// GetHandlers implements handlers.HandlerInterface.
func (h *productHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/products", HandlerFunc: h.get},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/products/:id", HandlerFunc: h.getProductById},
	)
	return
}
func NewProductHandler(
	path string,
	viewProductUsecase uc.ViewProductUsecase,
	searchProductByNameUsecase uc.SearchProductByNameUsecase, GetByIdProductUsecase uc.GetByIdProductUsecase) hd.HandlerInterface {
	return &productHandler{
		path:                       path,
		viewProductUsecase:         viewProductUsecase,
		searchProductByNameUsecase: searchProductByNameUsecase,
		getByIdProduct:             GetByIdProductUsecase,
	}
}
