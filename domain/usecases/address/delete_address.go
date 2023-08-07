package address

type DeleteAddressUsecase interface {
	DeleteAddressById(id uint64) (err error)
}
