package category

import (
	"github.com/phincon-backend/laza/domain/repositories/category"
	usecase "github.com/phincon-backend/laza/domain/usecases/category"
	repo "github.com/phincon-backend/laza/internal/repo/category"
)

type deleteCategoryByIdUsecaseImpl struct {
	deleteCategoryAction category.DeleteByIdAction
}

func (d deleteCategoryByIdUsecaseImpl) Execute(categoryId uint64) (rowAffected int64, err error) {
	rowAffected, err = d.deleteCategoryAction.DeleteById(categoryId)
	return
}

func NewDeleteCategoryByIdUsecaseImpl(categoryRepo repo.CategoryRepo) usecase.DeleteCategoryByIdUsecase {
	return &deleteCategoryByIdUsecaseImpl{
		deleteCategoryAction: &categoryRepo,
	}
}
