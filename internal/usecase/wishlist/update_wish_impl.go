package wishlist

import (
	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	wir "github.com/phincon-backend/laza/domain/repositories/wishlist"
	"github.com/phincon-backend/laza/domain/requests"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
)

type UpdateWishListUsecaseImpl struct {
	updateWishlist wir.UpdateWishListAction
	validate       *validator.Validate
}

func (u *UpdateWishListUsecaseImpl) Execute(userId uint64, rb requests.WishlistRequest) (value any, err error) {

	err = u.validate.Struct(rb)

	if err != nil {
		return
	}

	_md := model.Wishlist{UserId: userId, ProductId: rb.ProductId}

	value, err = u.updateWishlist.UpdateWishList(_md)

	return
}

func NewUpdateWishListUsecaseImpl(
	updateWishlist wir.UpdateWishListAction,
	validate *validator.Validate,

) p.UpdateWishListUsecase {
	return &UpdateWishListUsecaseImpl{
		updateWishlist: updateWishlist,
		validate:       validate,
	}
}
