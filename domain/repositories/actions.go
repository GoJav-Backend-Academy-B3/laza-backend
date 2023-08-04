package repositories

// EntityT = Entity Type

type GetAllAction[EntityT any] interface {
	GetAll() ([]EntityT, error)
}

type GetWithLimitAction[EntityT any] interface {
	GetWithLimit(offset, limit uint64) ([]EntityT, error)
}

type GetByIdAction[EntityT any] interface {
	GetById(id any) (EntityT, error)
}

type InsertAction[EntityT any] interface {
	Insert(dao EntityT) (EntityT, error)
}

type UpdateAction[EntityT any] interface {
	Update(id any, dao EntityT) (EntityT, error)
}

type DeleteAction[EntityT any] interface {
	Delete(id any) error
}

type BasicAction[EntityT any] interface {
	GetAllAction[EntityT]
	GetByIdAction[EntityT]
	InsertAction[EntityT]
	UpdateAction[EntityT]
	DeleteAction[EntityT]
}
