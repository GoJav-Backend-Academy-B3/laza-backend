package products

import (
	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
	"net/http"
)

type productHandler struct {
	path                       string
	createProductUsecase       uc.CreateProductUsecase
	updateProductUsecase       uc.UpdateProductUsecase
	viewProductUsecase         uc.ViewProductUsecase
	deleteProductUsecase       uc.DeleteProductUsecase
	searchProductByNameUsecase uc.SearchProductByNameUsecase
	getByIdProduct             uc.GetByIdProductUsecase
}

// GetHandlers implements handlers.HandlerInterface.

func (h *productHandler) GetHandlers() (hs []hd.HandlerStruct) {

	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodGet,
		Path:        h.path,
		HandlerFunc: h.get,
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodGet,
		Path:        h.path + "/:id",
		HandlerFunc: h.getProductById,
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
	searchProductByNameUsecase uc.SearchProductByNameUsecase,
	GetByIdProductUsecase uc.GetByIdProductUsecase,
) hd.HandlerInterface {
	return &productHandler{
		path:                       path,
		createProductUsecase:       createProductUsecase,
		updateProductUsecase:       updateProductUsecase,
		viewProductUsecase:         viewProductUsecase,
		deleteProductUsecase:       deleteProductUsecase,
		searchProductByNameUsecase: searchProductByNameUsecase,
		getByIdProduct:             GetByIdProductUsecase,
	}
}
