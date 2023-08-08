package wishlist

import (
	"errors"
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/product"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
	"gorm.io/gorm"

	"github.com/phincon-backend/laza/helper"
)

type UpdateWishListUsecaseImpl struct {
	updateWishlist       d.UpdateAction[m.Wishlist]
	getProductByIdAction action.GetProductByIdAction[m.Product]
}

func (u *UpdateWishListUsecaseImpl) Execute(userId, productId uint64) *helper.Response {
	_, err := u.getProductByIdAction.GetProductById(productId)

	if err == gorm.ErrRecordNotFound {
		return helper.GetResponse(errors.New("product not found").Error(), http.StatusNotFound, true)
	}

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	wishlist, err := u.updateWishlist.Update("", m.Wishlist{UserId: userId, ProductId: productId})

	return helper.GetResponse(wishlist, http.StatusOK, false)
}

func NewUpdateWishListUsecaseImpl(
	uw d.UpdateAction[m.Wishlist],
	gpi action.GetProductByIdAction[m.Product],
) p.UpdateWishListUsecase {
	return &UpdateWishListUsecaseImpl{
		updateWishlist:       uw,
		getProductByIdAction: gpi,
	}
}
