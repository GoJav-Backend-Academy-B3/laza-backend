package provider

import (
	"github.com/go-playground/validator/v10"
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
	newValidate := validator.New()

	wishlistRepo := r.NewWishList(gorm)

	wishlistUpdate := u.NewUpdateWishListUsecaseImpl(wishlistRepo, newValidate)
	wishlistGetLimit := u.NewgetWishlistLimitUsecase(wishlistRepo)
	return h.NewgetWishlistHandler(wishlistUpdate, wishlistGetLimit)

}
