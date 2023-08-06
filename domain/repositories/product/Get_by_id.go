package product

type GetProductByIdAction[EntityT any] interface {
	GetProductById(id any) (EntityT, error)
}
