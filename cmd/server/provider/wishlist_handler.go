package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/wishlist"

	u "github.com/phincon-backend/laza/internal/usecase/wishlist"

	r "github.com/phincon-backend/laza/internal/repo/wishlist"

	b "github.com/phincon-backend/laza/internal/db"
)

func NewWishListsHandler() d.HandlerInterface {

	// TODO: instantiate or get db
	db := b.GetPostgreSQLConnection()
	gorm := db.(*b.PsqlDB).Dbs

	repo := r.NewWishList(gorm)

	viewProduct := u.NewUpdateWishListUsecaseImpl(repo)
	return h.NewgetWishlistHandler(viewProduct)
}
