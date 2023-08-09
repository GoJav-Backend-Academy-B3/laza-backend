package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/products"

	u "github.com/phincon-backend/laza/internal/usecase/product"

	r "github.com/phincon-backend/laza/internal/repo/product"

	rs "github.com/phincon-backend/laza/internal/repo/size"

	b "github.com/phincon-backend/laza/internal/db"
	rc "github.com/phincon-backend/laza/internal/repo/category"
	rv "github.com/phincon-backend/laza/internal/repo/review"
)

func NewProductsHandler() d.HandlerInterface {

	// TODO: instantiate or get db
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	productRepo := r.NewProductRepo(gorm)
	sizeRepo := rs.NewSizeRepo(gorm)
	categoryRepo := rc.NewCategoryRepo(gorm)
	reviewRepo := rv.NewReviewRepo(gorm)
	// getSizeByIdUsecase := usecase.NewGetSizeByIdUsecaseImpl(sizeRepo)

	viewProduct := u.NewViewProductUsecaseImpl(productRepo)
	searchProduct := u.NewSearchProductUsecaseImpl(productRepo)
	getByIdProduct := u.NewGetByIdProductUsecase(productRepo, reviewRepo, sizeRepo)
	createProduct := u.NewCreateProductUsecaseImpl(productRepo, sizeRepo, categoryRepo)
	updateProduct := u.NewUpdateProductUsecaseImpl(productRepo, sizeRepo, categoryRepo)
	deleteProduct := u.NewDeleteProductUsecaseImpl(productRepo)
	return h.NewProductHandler("/products",
		createProduct, updateProduct, viewProduct, deleteProduct, searchProduct, getByIdProduct)
}
