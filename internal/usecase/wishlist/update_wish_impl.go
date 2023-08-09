package wishlist

import (
	"errors"
	"net/http"

	m "github.com/phincon-backend/laza/domain/model"
	d "github.com/phincon-backend/laza/domain/repositories"
	p "github.com/phincon-backend/laza/domain/usecases/wishlist"
	"gorm.io/gorm"

	"github.com/phincon-backend/laza/helper"
)

type UpdateWishListUsecaseImpl struct {
	updateWishlist       d.UpdateAction[m.Wishlist]
	getProductByIdAction d.GetByIdAction[m.Product]
}

func (u *UpdateWishListUsecaseImpl) Execute(userId, productId uint64) *helper.Response {
	_, err := u.getProductByIdAction.GetById(productId)

	if err == gorm.ErrRecordNotFound {
		return helper.GetResponse(errors.New("product not found").Error(), http.StatusNotFound, true)
	}

	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	wishlist, err := u.updateWishlist.Update("", m.Wishlist{UserId: userId, ProductId: productId})
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	if wishlist.ProductId == 0 && wishlist.UserId == 0 {
		return helper.GetResponse("successfully deleted wishlist", http.StatusOK, false)
	}

	return helper.GetResponse("successfully add wishlist", http.StatusOK, false)
}

func NewUpdateWishListUsecaseImpl(
	uw d.UpdateAction[m.Wishlist],
	gpi d.GetByIdAction[m.Product],
) p.UpdateWishListUsecase {
	return &UpdateWishListUsecaseImpl{
		updateWishlist:       uw,
		getProductByIdAction: gpi,
	}
}
