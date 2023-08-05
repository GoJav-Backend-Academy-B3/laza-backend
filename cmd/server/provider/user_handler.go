package provider

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/user"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"
	repoToken "github.com/phincon-backend/laza/internal/repo/verification_token"
	usecase "github.com/phincon-backend/laza/internal/usecase/user"
)

func NewUserHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	repoUser := repoUser.NewUserRepo(gorm)
	repoToken := repoToken.NewVerificationTokenRepo(gorm)

	getAllUser := usecase.NewGetAllUserUsecase(repoUser)
	getByIdUser := usecase.NewGetByIdUserUsecase(repoUser)
	getWithLimitUser := usecase.NewGetWithLimitUserUsecase(repoUser)
	insertUser := usecase.NewInsertUserUsecase(repoUser, repoToken, repoUser, repoUser)
	updateUser := usecase.NewUpdateUserUsecase(repoUser)
	deleteUser := usecase.NewDeleteUserUsecase(repoUser)

	return handler.NewUserHandler(getAllUser, getByIdUser, getWithLimitUser, insertUser, updateUser, deleteUser)
}
