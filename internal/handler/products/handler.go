package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/product"
	"github.com/phincon-backend/laza/middleware"
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
		HandlerFunc: h.getById,
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodPost,
		Path:        h.path,
		HandlerFunc: h.post,
		Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodPut,
		Path:        h.path + "/:id",
		HandlerFunc: h.put,
		Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
	})
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodDelete,
		Path:        h.path + "/:id",
		HandlerFunc: h.delete,
		Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
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
