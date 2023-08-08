package size

type DeleteSizeUsecase interface {
	Execute(id uint64) error
}
