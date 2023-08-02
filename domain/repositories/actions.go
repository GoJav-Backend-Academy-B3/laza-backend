package repositories

// EntityT = Entity Type

type GetAllAction[EntityT any] interface {
	GetAll() ([]EntityT, error)
}

type GetByIdAction[EntityT any] interface {
	GetById(id string) (EntityT, error)
}

type InsertAction[EntityT any] interface {
	Insert(dao EntityT) (string, error)
}

type UpdateAction[EntityT any] interface {
	Update(id string, dao EntityT) error
}

type DeleteAction[EntityT any] interface {
	Delete(id string) error
}

type BasicAction[EntityT any] interface {
	GetAllAction[EntityT]
	GetByIdAction[EntityT]
	InsertAction[EntityT]
	UpdateAction[EntityT]
	DeleteAction[EntityT]
}
