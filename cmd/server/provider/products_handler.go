package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/products"

	u "github.com/phincon-backend/laza/internal/usecase/product"

	br "github.com/phincon-backend/laza/internal/repo/brand"
	cr "github.com/phincon-backend/laza/internal/repo/category"
	pr "github.com/phincon-backend/laza/internal/repo/product"
	rr "github.com/phincon-backend/laza/internal/repo/review"
	sr "github.com/phincon-backend/laza/internal/repo/size"

	b "github.com/phincon-backend/laza/internal/db"
)

func NewProductsHandler() d.HandlerInterface {

	// TODO: instantiate or get db
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	productRepo := pr.NewProductRepo(gorm)
	sizeRepo := sr.NewSizeRepo(gorm)
	brandRepo := br.NewBrandRepo(gorm)
	categoryRepo := cr.NewCategoryRepo(gorm)
	reviewRepo := rr.NewReviewRepo(gorm)

	viewProduct := u.NewViewProductUsecaseImpl(productRepo)
	searchProduct := u.NewSearchProductUsecaseImpl(productRepo)
	getByIdProduct := u.NewGetByIdProductUsecase(productRepo, reviewRepo, sizeRepo, categoryRepo)
	createProduct := u.NewCreateProductUsecaseImpl(productRepo, brandRepo, sizeRepo, categoryRepo)
	updateProduct := u.NewUpdateProductUsecaseImpl(productRepo, brandRepo, sizeRepo, categoryRepo)
	deleteProduct := u.NewDeleteProductUsecaseImpl(productRepo)
	return h.NewProductHandler("/products",
		createProduct, updateProduct, viewProduct, deleteProduct, searchProduct, getByIdProduct)
}
