package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/products"

	u "github.com/phincon-backend/laza/internal/usecase/product"

	r "github.com/phincon-backend/laza/internal/repo/product"

	b "github.com/phincon-backend/laza/internal/db"
)

func NewProductsHandler() d.HandlerInterface {

	// TODO: instantiate or get db
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	repo := r.NewProductRepo(gorm)

	viewProduct := u.NewViewProductUsecaseImpl(repo)
	searchProduct := u.NewSearchProductUsecaseImpl(repo)
	return h.NewProductHandler("/products", viewProduct, searchProduct)
}
