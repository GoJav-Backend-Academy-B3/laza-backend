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

	cartrepo := repo.NewCartRepo(gorm)

	insertCartUc := usecase.NewinsertCartUsecase(cartrepo)
	deleteCartUc := usecase.NewdeleteCartUsecase(cartrepo)
	updateCartUc := usecase.NewupdateCartUsecase(cartrepo, cartrepo)
	getCartByIUc := usecase.NewgetCartByIdUsecase(cartrepo)

	return handler.NewcartHandler(insertCartUc, deleteCartUc, updateCartUc, getCartByIUc)

}
