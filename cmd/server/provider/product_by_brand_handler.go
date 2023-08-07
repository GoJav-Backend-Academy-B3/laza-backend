package provider

import (
	hd "github.com/phincon-backend/laza/domain/handlers"
	b "github.com/phincon-backend/laza/internal/db"
	handler "github.com/phincon-backend/laza/internal/handler/products"
	r "github.com/phincon-backend/laza/internal/repo/product"
	"github.com/phincon-backend/laza/internal/usecase/product"
)

func NewViewProductByBrandHandler() hd.HandlerInterface {
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	repo := r.NewProductRepo(gorm)

	viewByBrandUsecase := product.NewViewProductByBrandUsecaseImpl(repo)
	return handler.NewViewProductByBrandHandler("/products/brand", viewByBrandUsecase)
}
