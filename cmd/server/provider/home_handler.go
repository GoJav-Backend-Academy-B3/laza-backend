package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	handler "github.com/phincon-backend/laza/internal/handler/home"
)

func NewHomeHandler() domain.HandlerInterface {
	return handler.NewHomeHandler()
}
