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
	refreshTokenUser     auth.RefreshTokenUsecase

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
	refreshTokenUser auth.RefreshTokenUsecase,
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
		refreshTokenUser:     refreshTokenUser,
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
			Path:        "/auth/confirm",
			HandlerFunc: h.confirmEmail,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/confirm/resend",
			HandlerFunc: h.resendEmail,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/forgotpassword",
			HandlerFunc: h.forgotPassword,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/recover/code",
			HandlerFunc: h.verificationCode,
		},
		handlers.HandlerStruct{
			Method:      http.MethodPost,
			Path:        "/auth/recover/password",
			HandlerFunc: h.resetPassword,
		},
		handlers.HandlerStruct{
			Method:      http.MethodGet,
			Path:        "/auth/refresh",
			HandlerFunc: h.refreshToken,
		},
	)

	return
}
