package address

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/address"
	"github.com/phincon-backend/laza/middleware"
)

type addressHandler struct {
	basicPath string
	insert    address.AddAddressUsecase
	get       address.GetAddressUsecase
	update    address.UpdateAddressUsecase
	delete    address.DeleteAddressUsecase

	validate *validator.Validate
}

// GetHandlers implements handlers.HandlerInterface.
func (h *addressHandler) GetHandlers() []handlers.HandlerStruct {
	var hs []handlers.HandlerStruct

	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        h.basicPath,
			HandlerFunc: h.PostAddressHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basicPath,
			HandlerFunc: h.GetAllAddressByUserIdHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        h.basicPath + "/:id",
			HandlerFunc: h.GetAddressByIdHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        h.basicPath + "/:id",
			HandlerFunc: h.UpdateAddressHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        h.basicPath + "/:id",
			HandlerFunc: h.DeleteAddressHandler,
			Middlewares: []gin.HandlerFunc{middleware.AuthMiddleware()},
		},
	)

	return hs

}

func NewAddressHandler(
	basicPath string,
	insert address.AddAddressUsecase,
	get address.GetAddressUsecase,
	update address.UpdateAddressUsecase,
	delete address.DeleteAddressUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &addressHandler{
		basicPath: basicPath,
		insert:    insert,
		get:       get,
		update:    update,
		delete:    delete,
		validate:  validate,
	}
}
