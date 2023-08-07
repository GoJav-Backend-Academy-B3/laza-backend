package products

import (
	hd "github.com/phincon-backend/laza/domain/handlers"
	usecase "github.com/phincon-backend/laza/domain/usecases/product"
	"net/http"
)

type viewProductByBrandHandler struct {
	path                      string
	viewProductByBrandUsecase usecase.ViewProductByBrandUsecase
}

// GetHandlers implements handlers.HandlerInterface.
func (pb *viewProductByBrandHandler) GetHandlers() (hs []hd.HandlerStruct) {
	hs = append(hs, hd.HandlerStruct{
		Method:      http.MethodGet,
		Path:        pb.path,
		HandlerFunc: pb.get,
	})
	return
}

func NewViewProductByBrandHandler(path string, viewProductByBrandUC usecase.ViewProductByBrandUsecase) hd.HandlerInterface {
	return &viewProductByBrandHandler{
		path:                      path,
		viewProductByBrandUsecase: viewProductByBrandUC,
	}
}
