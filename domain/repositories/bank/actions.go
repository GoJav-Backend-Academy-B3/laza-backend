package bank

type GetAllBank[EntityT any] interface {
	GetAll() ([]EntityT, error)
}
type GetBankByIdAction[EntityT any] interface {
	GetBankById(id any) (EntityT, error)
}

type InsertBank[EntityT any] interface {
	Insert(dao EntityT) (EntityT, error)
}

type UpdateBank[EntityT any] interface {
	Update(id any, dao EntityT) (EntityT, error)
}

type DeleteBank[EntityT any] interface {
	Delete(id any) error
}
type ExistsBank interface {
	ExistsBank(name string) bool
}
