package request

type ProductOrder struct {
	Id       uint64 `json:"id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
