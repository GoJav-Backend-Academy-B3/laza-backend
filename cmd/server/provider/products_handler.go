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

	productRepo := r.NewProductRepo(gorm)
	sizeRepo := rs.NewSizeRepo(gorm)
	categoryRepo := rc.NewCategoryRepo(gorm)
	reviewRepo := rv.NewReviewRepo(gorm)
	brandRepo := br.NewBrandRepo(gorm)

	viewProduct := u.NewViewProductUsecaseImpl(productRepo)
	searchProduct := u.NewSearchProductUsecaseImpl(productRepo)
	getByIdProduct := u.NewGetByIdProductUsecase(productRepo, reviewRepo, sizeRepo, categoryRepo)
	createProduct := u.NewCreateProductUsecaseImpl(productRepo, sizeRepo, categoryRepo)
	createProduct := u.NewCreateProductUsecaseImpl(productRepo, brandRepo, sizeRepo, categoryRepo)
	updateProduct := u.NewUpdateProductUsecaseImpl(productRepo, sizeRepo, categoryRepo)
	deleteProduct := u.NewDeleteProductUsecaseImpl(productRepo)
	return h.NewProductHandler("/products",
		createProduct, updateProduct, viewProduct, deleteProduct, searchProduct, getByIdProduct)
}
