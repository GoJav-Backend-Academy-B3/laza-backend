package bank

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/bank"
)

type bankHandler struct {
	getAllBank  bank.GetAllBankUsecase
	getByIdBank bank.GetByIdBankUsecase
	insertBank  bank.InsertBankUsecase
	updateBank  bank.UpdateBankUsecase
	deleteBank  bank.DeleteBankUsecase
}

func NewBankHandler(
	getAllBank bank.GetAllBankUsecase,
	getByIdBank bank.GetByIdBankUsecase,
	insertBank bank.InsertBankUsecase,
	updateBank bank.UpdateBankUsecase,
	deleteBank bank.DeleteBankUsecase,
) handlers.HandlerInterface {
	return &bankHandler{
		getAllBank:  getAllBank,
		getByIdBank: getByIdBank,
		insertBank:  insertBank,
		updateBank:  updateBank,
		deleteBank:  deleteBank,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *bankHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/bank", HandlerFunc: h.get},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/bank/:id", HandlerFunc: h.getBankById},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/bank", HandlerFunc: h.insert},
		handlers.HandlerStruct{Method: http.MethodPut, Path: "/bank/:id", HandlerFunc: h.update},
		handlers.HandlerStruct{Method: http.MethodDelete, Path: "/bank/:id", HandlerFunc: h.delete},
		// handlers.HandlerStruct{Method: http.MethodGet, Path: "/user/", HandlerFunc: h.getWithLimit},
	)
	return
}