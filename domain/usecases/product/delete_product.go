package product

type DeleteProductUsecase interface {
	Execute(id uint64) error
}
