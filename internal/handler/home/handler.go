package handler

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
)

type HomeHandler struct {
}

// GetHandlers implements handlers.HandlerInterface.
func (h *HomeHandler) GetHandlers() (ar []handlers.HandlerStruct) {
	ar = append(ar, handlers.HandlerStruct{
		Method:      http.MethodGet,
		Path:        "/",
		HandlerFunc: h.get,
	})

	return
}

func NewHomeHandler() handlers.HandlerInterface {
	return &HomeHandler{}
}
