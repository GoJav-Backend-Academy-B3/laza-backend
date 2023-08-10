package brand

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/brand"
	"github.com/phincon-backend/laza/middleware"
)

type brandHandler struct {
	basePath                   string
	createBrandUsecase         brand.CreateBrandUsecase
	searchBrandByNameUsecase   brand.SearchBrandByNameUsecase
	deleteBrandUsecase         brand.DeleteBrandByIdUsecase
	getBrandByIdUsecase        brand.GetBrandByIdUsecase
	updateBrandNameByIdUsecase brand.UpdateBrandNameByIdUsecase
	viewBrandUsecase           brand.ViewBrandUsecase
	validate                   *validator.Validate
}

// GetHandlers implements handlers.HandlerInterface.
func (h *brandHandler) GetHandlers() []handlers.HandlerStruct {
	var hs []handlers.HandlerStruct
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        h.basePath,
			HandlerFunc: h.PostBrandHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath + "/search",
			HandlerFunc: h.SearchByBrandName,
			Middlewares: []gin.HandlerFunc{},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath + "/:id",
			HandlerFunc: h.GetById,
			Middlewares: []gin.HandlerFunc{},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basePath,
			HandlerFunc: h.ViewAllBrand,
			Middlewares: []gin.HandlerFunc{},
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        h.basePath + "/:id",
			HandlerFunc: h.UpdateBrand,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        h.basePath + "/:id",
			HandlerFunc: h.DeleteBrandById,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
	)

	return hs
}

func NewBrandHandler(
	basePath string,
	createBrandUsecase brand.CreateBrandUsecase,
	searchBrandByNameUsecase brand.SearchBrandByNameUsecase,
	deleteBrandUsecase brand.DeleteBrandByIdUsecase,
	getBrandByIdUsecase brand.GetBrandByIdUsecase,
	updateBrandNameByIdUsecase brand.UpdateBrandNameByIdUsecase,
	viewBrandUsecase brand.ViewBrandUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &brandHandler{
		basePath:                   basePath,
		createBrandUsecase:         createBrandUsecase,
		searchBrandByNameUsecase:   searchBrandByNameUsecase,
		getBrandByIdUsecase:        getBrandByIdUsecase,
		updateBrandNameByIdUsecase: updateBrandNameByIdUsecase,
		deleteBrandUsecase:         deleteBrandUsecase,
		viewBrandUsecase:           viewBrandUsecase,
		validate:                   validate,
	}
}
