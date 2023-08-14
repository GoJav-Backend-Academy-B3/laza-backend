package provider

import (
	"github.com/go-playground/validator/v10"
	domain "github.com/phincon-backend/laza/domain/handlers"
	"github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/cart"
	repo "github.com/phincon-backend/laza/internal/repo/cart"

	usecase "github.com/phincon-backend/laza/internal/usecase/cart"
)

func NewCartHandler() domain.HandlerInterface {

	dbs := db.GetPostgreSQLConnection()
	gorm := dbs.(*db.PsqlDB).Dbs

	validator := validator.New()

	cartrepo := repo.NewCartRepo(gorm)

	insertCartUc := usecase.NewinsertCartUsecase(cartrepo, validator)
	deleteCartUc := usecase.NewdeleteCartUsecase(cartrepo, validator)
	updateCartUc := usecase.NewupdateCartUsecase(cartrepo, validator)
	getCartByIUc := usecase.NewgetCartByIdUsecase(cartrepo, cartrepo)

	return handler.NewcartHandler(insertCartUc, deleteCartUc, updateCartUc, getCartByIUc)

}
