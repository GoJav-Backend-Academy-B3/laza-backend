package category

type DeleteByIdAction interface {
	DeleteById(id uint64) (rowAffected int64, err error)
}
