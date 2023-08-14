package bank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/middleware"
)

type bankHandler struct {
	getAllBank  bank.GetAllBankUsecase
	getByIdBank bank.GetByIdBankUsecase
	insertBank  bank.InsertBankUsecase
	updateBank  bank.UpdateBankUsecase
	deleteBank  bank.DeleteBankUsecase
	validate    *validator.Validate
}

func NewBankHandler(
	getAllBank bank.GetAllBankUsecase,
	getByIdBank bank.GetByIdBankUsecase,
	insertBank bank.InsertBankUsecase,
	updateBank bank.UpdateBankUsecase,
	deleteBank bank.DeleteBankUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &bankHandler{
		getAllBank:  getAllBank,
		getByIdBank: getByIdBank,
		insertBank:  insertBank,
		updateBank:  updateBank,
		deleteBank:  deleteBank,
		validate:    validate,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *bankHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/bank",
			HandlerFunc: h.get,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/bank/:id",
			HandlerFunc: h.getBankById,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/bank",
			HandlerFunc: h.insert,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/bank/:id",
			HandlerFunc: h.update,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        "/bank/:id",
			HandlerFunc: h.delete,
			Middlewares: gin.HandlersChain{middleware.AuthMiddleware(), middleware.AdminRoleMiddleware()},
		},
	)
	return
}
