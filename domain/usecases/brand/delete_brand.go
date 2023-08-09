package brand

type DeleteBrandByIdUsecase interface {
	Execute(brandId uint64) (err error)
}
