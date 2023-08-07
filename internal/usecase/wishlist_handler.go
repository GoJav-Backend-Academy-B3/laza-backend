package provider

// import (
// 	d "github.com/phincon-backend/laza/domain/handlers"
// 	h "github.com/phincon-backend/laza/internal/handler/wishlist"

// 	u "github.com/phincon-backend/laza/internal/usecase/wishlist"

// 	b "github.com/phincon-backend/laza/internal/db"
// )

// func NewWishListsHandler() d.HandlerInterface {

// 	// TODO: instantiate or get db
// 	db := b.GetPostgreSQLConnection()
// 	gorm := db.(*b.PsqlDB).Dbs

// 	repo := r.NewWishList(gorm)

// 	viewProduct := u.NewUpdateWishListUsecaseImpl(repo)
// 	return h.NewgetWishlistHandler("/wishlist/:userId/:productId", viewProduct)
// }

// wishlistRepo := r.NewWishList(gorm)
// productRepo := p.NewProductRepo(gorm)

// wishlistUpdate := u.NewUpdateWishListUsecaseImpl(wishlistRepo)
// wistlistGet := u.NewgetWishlistUsecase(wishlistRepo, productRepo)
// return h.NewgetWishlistHandler(wishlistUpdate, wistlistGet)
