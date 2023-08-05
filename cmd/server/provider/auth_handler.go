package provider

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/auth"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"
	repoToken "github.com/phincon-backend/laza/internal/repo/verification_token"
	usecaseAuth "github.com/phincon-backend/laza/internal/usecase/auth"
	usecaseUser "github.com/phincon-backend/laza/internal/usecase/user"
)

func NewAuthHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	repoUser := repoUser.NewUserRepo(gorm)
	repoToken := repoToken.NewVerificationTokenRepo(gorm)

	loginUser := usecaseAuth.NewLoginUserUsecase(repoUser)
	registerUser := usecaseUser.NewInsertUserUsecase(repoUser, repoToken, repoUser, repoUser)
	verifyEmailUser := usecaseAuth.NewVerifyEmailUserUsecase(repoUser, repoUser, repoToken)
	resendEmailUser := usecaseAuth.NewResendEmailUserUsecase(repoToken, repoUser, repoUser)
	
	return handler.NewAuthHandler(loginUser, registerUser, verifyEmailUser, resendEmailUser)
}
