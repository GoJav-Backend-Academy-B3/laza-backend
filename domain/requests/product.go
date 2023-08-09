package requests

type ProductOrder struct {
	Id       uint64 `json:"id,omitempty" binding:"required"`
	Quantity int    `json:"quantity,omitempty" binding:"required"`
}
