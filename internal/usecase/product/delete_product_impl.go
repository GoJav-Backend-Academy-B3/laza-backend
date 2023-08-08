package product

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	usecase "github.com/phincon-backend/laza/domain/usecases/product"
	"gorm.io/gorm"
)

type DeleteProductUsecaseImpl struct {

	// Delete product action
	deleteProductAction repositories.DeleteAction[model.Product]
}

// Execute implements product.DeleteProductUsecase.
func (u *DeleteProductUsecaseImpl) Execute(id uint64) error {

	err := u.deleteProductAction.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("NotFound: Id tidak ditemukan")
	}
	return nil
}

func NewDeleteProductUsecaseImpl(
	deleteAction repositories.DeleteAction[model.Product],
) usecase.DeleteProductUsecase {
	return &DeleteProductUsecaseImpl{
		deleteProductAction: deleteAction,
	}
}
