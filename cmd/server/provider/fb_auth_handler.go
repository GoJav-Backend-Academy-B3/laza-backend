package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"

	handler "github.com/phincon-backend/laza/internal/handler/facebook_auth"

	uc "github.com/phincon-backend/laza/internal/usecase/facebook_auth"
)

func NewFacebookAuthHandler() domain.HandlerInterface {
	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	repoUsers := repoUser.NewUserRepo(gorm)
	usecase := uc.NewFacebookAuthUsecase(*repoUsers)
	return handler.NewFacebookAuthHandler("/auth/facebook", "/auth/facebook/callback", usecase)
}
