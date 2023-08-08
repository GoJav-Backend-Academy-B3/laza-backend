package products

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
)

type productHandler struct {
	path                       string
	createProductUsecase       uc.CreateProductUsecase
	updateProductUsecase       uc.UpdateProductUsecase
	viewProductUsecase         uc.ViewProductUsecase
	deleteProductUsecase       uc.DeleteProductUsecase
	searchProductByNameUsecase uc.SearchProductByNameUsecase
}

// GetHandlers implements handlers.HandlerInterface.
func (h *productHandler) GetHandlers() (hs []hd.HandlerStruct) {
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodGet,
		Path:        h.path,
		HandlerFunc: h.get,
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodPost,
		Path:        h.path,
		HandlerFunc: h.post,
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodPut,
		Path:        h.path + "/:id",
		HandlerFunc: h.put,
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodDelete,
		Path:        h.path + "/:id",
		HandlerFunc: h.delete,
	})
	return
}

func NewProductHandler(
	path string,
	createProductUsecase uc.CreateProductUsecase,
	updateProductUsecase uc.UpdateProductUsecase,
	viewProductUsecase uc.ViewProductUsecase,
	deleteProductUsecase uc.DeleteProductUsecase,
	searchProductByNameUsecase uc.SearchProductByNameUsecase) hd.HandlerInterface {
	return &productHandler{
		path:                       path,
		createProductUsecase:       createProductUsecase,
		updateProductUsecase:       updateProductUsecase,
		viewProductUsecase:         viewProductUsecase,
		deleteProductUsecase:       deleteProductUsecase,
		searchProductByNameUsecase: searchProductByNameUsecase,
	}
}
