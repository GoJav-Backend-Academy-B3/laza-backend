package auth

import (
	"net/http"

	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type authHandler struct {
	loginUser       auth.LoginUserUsecase
	registerUser    user.InsertUserUsecase
	verifyEmailUser auth.VerifyEmailUserUsecase
	resendEmailUser auth.ResendEmailUserUsecase
}

func NewAuthHandler(loginUser auth.LoginUserUsecase, registerUser user.InsertUserUsecase,
	verifyEmailUser auth.VerifyEmailUserUsecase, resendEmailUser auth.ResendEmailUserUsecase) handlers.HandlerInterface {
	return &authHandler{
		loginUser:       loginUser,
		registerUser:    registerUser,
		verifyEmailUser: verifyEmailUser,
		resendEmailUser: resendEmailUser,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *authHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/login", HandlerFunc: h.login},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/register", HandlerFunc: h.register},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/auth/resend_email", HandlerFunc: h.resendEmail},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/auth/confirm_email/", HandlerFunc: h.verifyEmail},
	)

	return
}
