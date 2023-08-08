package category

import (
	"github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/category"
	"net/http"
)

type categoryHandler struct {
	basePath                      string
	createCategoryUsecase         uc.CreateCategoryUsecase
	deleteCategoryByIdUsecase     uc.DeleteCategoryByIdUsecase
	getCategoryByIdUsecase        uc.GetCategoryByIdUsecase
	searchCategoryByNameUsecase   uc.SearchCategoryByNameUsecase
	updateCategoryNameByIdUsecase uc.UpdateCategoryNameByIdUsecase
	viewCategoryUsecase           uc.ViewCategoryUsecase
}

func (ch *categoryHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        ch.basePath,
			HandlerFunc: ch.postCategory,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        ch.basePath + "/:id",
			HandlerFunc: ch.getById,
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        ch.basePath + "/:id",
			HandlerFunc: ch.deleteById,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        ch.basePath + "/search", // search?name={name}
			HandlerFunc: ch.searchByName,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        ch.basePath,
			HandlerFunc: ch.updateNameById,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        ch.basePath + "/all",
			HandlerFunc: ch.getAll,
		},
	)
	return
}

func NewCategoryHandler(basePath string,
	createCategoryUsecase uc.CreateCategoryUsecase,
	deleteCategoryByIdUsecase uc.DeleteCategoryByIdUsecase,
	getCategoryByIdUsecase uc.GetCategoryByIdUsecase,
	searchCategoryByNameUsecase uc.SearchCategoryByNameUsecase,
	updateCategoryNameByIdUsecase uc.UpdateCategoryNameByIdUsecase,
	viewCategoryUsecase uc.ViewCategoryUsecase) handlers.HandlerInterface {
	return &categoryHandler{
		basePath:                      basePath,
		createCategoryUsecase:         createCategoryUsecase,
		deleteCategoryByIdUsecase:     deleteCategoryByIdUsecase,
		getCategoryByIdUsecase:        getCategoryByIdUsecase,
		searchCategoryByNameUsecase:   searchCategoryByNameUsecase,
		updateCategoryNameByIdUsecase: updateCategoryNameByIdUsecase,
		viewCategoryUsecase:           viewCategoryUsecase}
}
