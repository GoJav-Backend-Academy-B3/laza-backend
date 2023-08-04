package provider

import domain "github.com/phincon-backend/laza/domain/handlers"
import handler "github.com/phincon-backend/laza/internal/handler/facebook_auth"
import uc "github.com/phincon-backend/laza/internal/usecase/facebook_auth"

func NewFacebookAuthHandler() domain.HandlerInterface {
	usecase := uc.NewFacebookAuthUsecase()
	return handler.NewFacebookAuthHandler("/auth/facebook", "/auth/facebook/callback", usecase)
}
