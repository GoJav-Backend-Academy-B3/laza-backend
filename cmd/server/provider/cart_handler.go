package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/cart"
	repo "github.com/phincon-backend/laza/internal/repo/cart"
	usecase "github.com/phincon-backend/laza/internal/usecase/cart"
)

func NewCartHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	insertCartrepo := repo.NewCartRepo(gorm)

	insertCartUc := usecase.NewinsertCartUsecase(insertCartrepo)

	return handler.NewcartHandler(insertCartUc)

}
