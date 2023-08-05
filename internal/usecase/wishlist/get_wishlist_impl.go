package wishlist

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
)

type getWishlistUsecase struct {
	getWishlistRepo   repositories.GetByIdAction[model.Wishlist]
	getAllProductRepo repositories.GetAllAction[model.Product]
}
