package request

type CategoryRequest struct {
	Id       uint64 `json:"id"`
	Category string `json:"category" validate:"required"`
}

func (cr *CategoryRequest) GetCategory() string {
	return cr.Category
}
