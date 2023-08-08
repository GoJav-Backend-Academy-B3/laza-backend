package size

import (
	"net/http"

	hd "github.com/phincon-backend/laza/domain/handlers"
	uc "github.com/phincon-backend/laza/domain/usecases/size"
)

type sizeHandler struct {
	path               string
	addSizeUsecase     uc.AddSizeUsecase
	deleteSizeUsecase  uc.DeleteSizeUsecase
	getSizeByIdUsecase uc.GetSizeById
	getAllSizeUsecase  uc.GetAllSizeUsecase
	updateSizeUsecase  uc.UpdateSizeUsecase
}

// GetHandlers implements handlers.HandlerInterface.
func (h *sizeHandler) GetHandlers() (hs []hd.HandlerStruct) {
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

func NewSizeHandler(
	path string,
	addSizeUsecase uc.AddSizeUsecase,
	deleteSizeUsecase uc.DeleteSizeUsecase,
	getSizeByIdUsecase uc.GetSizeById,
	getAllSizeUsecase uc.GetAllSizeUsecase,
	updateSizeUsecase uc.UpdateSizeUsecase,
) hd.HandlerInterface {
	return &sizeHandler{
		path:               path,
		addSizeUsecase:     addSizeUsecase,
		deleteSizeUsecase:  deleteSizeUsecase,
		getSizeByIdUsecase: getSizeByIdUsecase,
		getAllSizeUsecase:  getAllSizeUsecase,
		updateSizeUsecase:  updateSizeUsecase,
	}
}
