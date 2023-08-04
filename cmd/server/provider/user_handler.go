package provider

import (
	"github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	hdle "github.com/phincon-backend/laza/internal/handler/user"
	repo "github.com/phincon-backend/laza/internal/repo/user"
	usca "github.com/phincon-backend/laza/internal/usecase/user"
)

func NewUserHandler() handlers.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	repo := repo.NewUserRepo(gorm)
	getAllUser := usca.NewGetAllUserUsecase(repo)
	getByIdUser := usca.NewGetByIdUserUsecase(repo)
	getWithLimitUser := usca.NewGetWithLimitUserUsecase(repo)
	insertUser := usca.NewInsertUserUsecase(repo)
	updateUser := usca.NewUpdateUserUsecase(repo)
	deleteUser := usca.NewDeleteUserUsecase(repo)

	return hdle.NewUserHandler(getAllUser, getByIdUser, getWithLimitUser, insertUser, updateUser, deleteUser)
}
