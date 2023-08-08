package provider

import (
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/cart"
	repo "github.com/phincon-backend/laza/internal/repo/cart"
	product_repo "github.com/phincon-backend/laza/internal/repo/product"

	usecase "github.com/phincon-backend/laza/internal/usecase/cart"
)

func NewCartHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	cartrepo := repo.NewCartRepo(gorm)
	productRepo := product_repo.NewProductRepo(gorm)

	insertCartUc := usecase.NewinsertCartUsecase(cartrepo, productRepo)
	deleteCartUc := usecase.NewdeleteCartUsecase(cartrepo, cartrepo)
	updateCartUc := usecase.NewupdateCartUsecase(cartrepo, cartrepo)
	getCartByIUc := usecase.NewgetCartByIdUsecase(cartrepo, cartrepo)

	return handler.NewcartHandler(insertCartUc, deleteCartUc, updateCartUc, getCartByIUc)

}
