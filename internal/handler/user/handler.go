package user

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type userHandler struct {
	getAllUser user.GetAllUserUsecase
}

// GetHandlers implements handlers.HandlerInterface.

func NewUserHandler(
	getAllUser user.GetAllUserUsecase,
) handlers.HandlerInterface {
	return &userHandler{
		getAllUser: getAllUser,
	}
}

func (h *userHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/user", HandlerFunc: h.get},
	)
	return
}
