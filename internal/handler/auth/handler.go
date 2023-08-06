package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type authHandler struct {
	loginUser          auth.LoginUserUsecase
	registerUser       user.InsertUserUsecase
	verifyEmailUser    auth.VerifyEmailUserUsecase
	resendEmailUser    auth.ResendEmailUserUsecase
	forgetPasswordUser auth.ForgetPasswordUserUsecase
	updatePasswordUser auth.UpdatePasswordUserUsecase

	validate *validator.Validate
}

func NewAuthHandler(
	loginUser auth.LoginUserUsecase,
	registerUser user.InsertUserUsecase,
	verifyEmailUser auth.VerifyEmailUserUsecase,
	resendEmailUser auth.ResendEmailUserUsecase,
	forgetPasswordUser auth.ForgetPasswordUserUsecase,
	updatePasswordUser auth.UpdatePasswordUserUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &authHandler{
		loginUser:          loginUser,
		registerUser:       registerUser,
		verifyEmailUser:    verifyEmailUser,
		resendEmailUser:    resendEmailUser,
		forgetPasswordUser: forgetPasswordUser,
		updatePasswordUser: updatePasswordUser,
		validate:           validate,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *authHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/login", HandlerFunc: h.login},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/register", HandlerFunc: h.register},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/auth/resend-verify", HandlerFunc: h.resendEmail},
		handlers.HandlerStruct{Method: http.MethodGet, Path: "/auth/verify-email/", HandlerFunc: h.verifyEmail},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/auth/forget-password", HandlerFunc: h.forgetPassword},
		handlers.HandlerStruct{Method: http.MethodPost, Path: "/auth/update-password/", HandlerFunc: h.updatePassword},
	)

	return
}
