package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/auth"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"
	repoCode "github.com/phincon-backend/laza/internal/repo/verification_code"
	repoToken "github.com/phincon-backend/laza/internal/repo/verification_token"
	usecaseAuth "github.com/phincon-backend/laza/internal/usecase/auth"
	usecaseUser "github.com/phincon-backend/laza/internal/usecase/user"
)

func NewAuthHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	validate := validator.New()

	repoUser := repoUser.NewUserRepo(gorm)
	repoToken := repoToken.NewVerificationTokenRepo(gorm)
	repoCode := repoCode.NewVerificationCodeRepo(gorm)

	loginUser := usecaseAuth.NewLoginUserUsecase(*repoUser)
	loginGoogleUser := usecaseAuth.NewLoginGoogleUserUsecase(*repoUser)
	registerUser := usecaseUser.NewInsertUserUsecase(*repoUser, *repoToken)
	verifyEmailUser := usecaseAuth.NewVerifyEmailUserUsecase(*repoUser, *repoToken)
	resendEmailUser := usecaseAuth.NewResendEmailUserUsecase(*repoUser, *repoToken)
	forgetPasswordUser := usecaseAuth.NewForgetPasswordUserUsecase(*repoUser, *repoCode)
	resetPasswordUser := usecaseAuth.NewResetPasswordUserUsecase(*repoUser, *repoCode)

	return handler.NewAuthHandler(loginUser, loginGoogleUser, registerUser, verifyEmailUser, resendEmailUser, forgetPasswordUser, resetPasswordUser, validate)
}
