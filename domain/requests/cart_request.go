package requests

type CartRequest struct {
	ProductId uint64 `json:"product_id" validate:"required"`
	SizeId    uint64 `json:"size_id" validate:"required"`
}
