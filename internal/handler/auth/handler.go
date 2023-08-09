package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/domain/usecases/auth"
	"github.com/phincon-backend/laza/domain/usecases/user"
)

type authHandler struct {
	loginUser            auth.LoginUserUsecase
	loginGoogleUser      auth.LoginGoogleUserUsecase
	registerUser         user.InsertUserUsecase
	verifyEmailUser      auth.VerifyEmailUserUsecase
	resendEmailUser      auth.ResendEmailUserUsecase
	forgotPasswordUser   auth.ForgotPasswordUserUsecase
	verificationCodeUser auth.VerificationCodeUserUsecase
	resetPasswordUser    auth.ResetPasswordUserUsecase

	validate *validator.Validate
}

func NewAuthHandler(
	loginUser auth.LoginUserUsecase,
	loginGoogleUser auth.LoginGoogleUserUsecase,
	registerUser user.InsertUserUsecase,
	verifyEmailUser auth.VerifyEmailUserUsecase,
	resendEmailUser auth.ResendEmailUserUsecase,
	forgotPasswordUser auth.ForgotPasswordUserUsecase,
	verificationCodeUser auth.VerificationCodeUserUsecase,
	resetPasswordUser auth.ResetPasswordUserUsecase,
	validate *validator.Validate,
) handlers.HandlerInterface {
	return &authHandler{
		loginUser:            loginUser,
		loginGoogleUser:      loginGoogleUser,
		registerUser:         registerUser,
		verifyEmailUser:      verifyEmailUser,
		resendEmailUser:      resendEmailUser,
		forgotPasswordUser:   forgotPasswordUser,
		verificationCodeUser: verificationCodeUser,
		resetPasswordUser:    resetPasswordUser,
		validate:             validate,
	}
}

// GetHandlers implements handlers.HandlerInterface.
func (h *authHandler) GetHandlers() (hs []handlers.HandlerStruct) {
	hs = append(hs,
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/login",
			HandlerFunc: h.login,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/register",
			HandlerFunc: h.register,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/auth/google",
			HandlerFunc: h.loginGoogle,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/auth/google/callback",
			HandlerFunc: h.loginGoogleCallback,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/auth/verify-email/",
			HandlerFunc: h.verifyEmail,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/resend-verify",
			HandlerFunc: h.resendEmail,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/forgot-password",
			HandlerFunc: h.forgotPassword,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/verification-code",
			HandlerFunc: h.verificationCode,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/reset-password/",
			HandlerFunc: h.resetPassword,
		},
	)

	return
}
