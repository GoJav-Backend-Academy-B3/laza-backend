package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/twitter_auth"
	repoUser "github.com/phincon-backend/laza/internal/repo/user"

	u "github.com/phincon-backend/laza/internal/usecase/twitter_auth"

	b "github.com/phincon-backend/laza/internal/db"
)

func NewtwitterAuthHandler() d.HandlerInterface {
	// TODO: instantiate or get db
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	repoUsers := repoUser.NewUserRepo(gorm)
	useCaseTwitter := u.NewtwitterAuthUsecase(repoUsers, repoUsers, repoUsers, repoUsers)

	return h.NewtwitterAuthHandler("/auth", "/auth/twitter/callback", useCaseTwitter)
}
