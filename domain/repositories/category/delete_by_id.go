package category

type DeleteByIdAction interface {
	DeleteById(id uint64) (err error)
}
