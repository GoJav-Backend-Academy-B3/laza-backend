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


	loginUser := usecaseAuth.NewLoginUserUsecase(repoUser)
	registerUser := usecaseUser.NewInsertUserUsecase(repoUser, repoToken, repoUser, repoUser)
	verifyEmailUser := usecaseAuth.NewVerifyEmailUserUsecase(repoUser, repoUser, repoToken)
	resendEmailUser := usecaseAuth.NewResendEmailUserUsecase(repoToken, repoUser, repoUser)

	return handler.NewAuthHandler(loginUser, registerUser, verifyEmailUser, resendEmailUser, validate)

	repoCode := repoCode.NewVerificationCodeRepo(gorm)

	loginUser := usecaseAuth.NewLoginUserUsecase(repoUser)
	loginGoogleUser := usecaseAuth.NewLoginGoogleUserUsecase(repoUser, repoUser, repoUser, repoUser)
	registerUser := usecaseUser.NewInsertUserUsecase(repoUser, repoToken, repoUser, repoUser)
	verifyEmailUser := usecaseAuth.NewVerifyEmailUserUsecase(repoUser, repoUser, repoToken)
	resendEmailUser := usecaseAuth.NewResendEmailUserUsecase(repoToken, repoUser, repoUser)
	forgetPasswordUser := usecaseAuth.NewForgetPasswordUserUsecase(repoCode, repoCode, repoCode, repoUser, repoUser)
	updatePasswordUser := usecaseAuth.NewUpdatePasswordUserUsecase(repoUser, repoUser, repoCode)

	return handler.NewAuthHandler(loginUser, loginGoogleUser, registerUser, verifyEmailUser, resendEmailUser, forgetPasswordUser, updatePasswordUser, validate)

}
