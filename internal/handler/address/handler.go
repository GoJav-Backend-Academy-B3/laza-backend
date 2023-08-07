package address

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type addressHandler struct {
	insert address.AddAddressUsecase
	get    address.GetAddressUsecase
	update address.UpdateAddressUsecase
	delete address.DeleteAddressUsecase

	validate *validator.Validate
}

// GetHandlers implements handlers.HandlerInterface.
func (h *addressHandler) GetHandlers() []handlers.HandlerStruct {
	var hs []handlers.HandlerStruct

	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/address",
			HandlerFunc: h.PostAddressHandler,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/address/user/:userId",
			HandlerFunc: h.GetAllAddressByUserIdHandler,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/address/:id",
			HandlerFunc: h.GetAddressByIdHandler,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPut,
			Path:        "/address/:id",
			HandlerFunc: h.UpdateAddressHandler,
		},
		handlers.HandlerStruct{
			Method:      http.MethodDelete,
			Path:        "/address/:id",
			HandlerFunc: h.DeleteAddressHandler,
		},
	)

	return hs

}

func NewAddressHandler(insert address.AddAddressUsecase,
	get address.GetAddressUsecase,
	update address.UpdateAddressUsecase,
	delete address.DeleteAddressUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &addressHandler{
		insert:   insert,
		get:      get,
		update:   update,
		delete:   delete,
		validate: validate,
	}
}
