package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/products"

	u "github.com/phincon-backend/laza/internal/usecase/product"

	r "github.com/phincon-backend/laza/internal/repo/product"
)

func NewProductsHandler() d.HandlerInterface {

	// TODO: instantiate or get db

	repo := r.NewProductRepo(db)

	viewProduct := u.NewViewProductUsecaseImpl(repo)
	return h.NewProductHandler("/products", viewProduct)
}
