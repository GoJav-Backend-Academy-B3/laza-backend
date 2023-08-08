package category

type DeleteCategoryByIdUsecase interface {
	Execute(categoryId uint64) (err error)
}
