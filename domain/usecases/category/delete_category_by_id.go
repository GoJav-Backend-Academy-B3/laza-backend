package category

type DeleteCategoryByIdUsecase interface {
	Execute(categoryId uint64) (rowAffected int64, err error)
}
