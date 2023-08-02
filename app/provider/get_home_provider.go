package provider

import (
	"github.com/phincon-backend/laza/domain/contract"
	"github.com/phincon-backend/laza/internal/delivery/handler"
)

func NewHomeHandler() contract.MainHandlerInterface {
	return handler.NewHomeHandler()
}
