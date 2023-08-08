package provider

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/user"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"
	usecase "github.com/phincon-backend/laza/internal/usecase/user"
)

func NewUserHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	validate := validator.New()

	repoUser := repoUser.NewUserRepo(gorm)

	getAllUser := usecase.NewGetAllUserUsecase(repoUser)
	getByIdUser := usecase.NewGetByIdUserUsecase(repoUser)
	getWithLimitUser := usecase.NewGetWithLimitUserUsecase(repoUser)
	updateUser := usecase.NewUpdateUserUsecase(repoUser, repoUser, repoUser, repoUser)
	deleteUser := usecase.NewDeleteUserUsecase(repoUser)
	changePasswordUser := usecase.NewChangePasswordUserUsecase(repoUser, repoUser)

	return handler.NewUserHandler(getAllUser, getByIdUser, getWithLimitUser, updateUser, changePasswordUser, deleteUser, validate)
}
