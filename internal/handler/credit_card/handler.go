package credit_card

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type getCreditCardHandler struct {
	addPath     string
	updatePath  string
	getByIdPath string
	getByAll    string
	addCcUc     uc.AddCreditCardUsecase
	updateCcUc  uc.UpdateCreditCardUsecase
	getByIdCcUc uc.GetByIdCreditCardUsecase
	getAllCcUc  uc.GetAllCreditCardUsecase
}

func (h *getCreditCardHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        h.addPath,
			HandlerFunc: h.Add,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        h.updatePath,
			HandlerFunc: h.Update,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.getByIdPath,
			HandlerFunc: h.GetById,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.getByAll,
			HandlerFunc: h.GetAll,
		},
	)
	return hs
}

func NewgetCreditCardHandler(
	addPath string,
	updatePath string,
	getByIdPath string,
	getByAll string,
	addCcUc uc.AddCreditCardUsecase,
	updateCcUc uc.UpdateCreditCardUsecase,
	getByIdCcUc uc.GetByIdCreditCardUsecase,
	getAllCcUc uc.GetAllCreditCardUsecase,

) handlers.HandlerInterface {
	return &getCreditCardHandler{
		addPath:     addPath,
		updatePath:  updatePath,
		getByIdPath: getByIdPath,
		getByAll:    getByAll,
		addCcUc:     addCcUc,
		updateCcUc:  updateCcUc,
		getByIdCcUc: getByIdCcUc,
		getAllCcUc:  getAllCcUc,
	}
}
