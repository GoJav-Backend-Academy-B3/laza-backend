package requests

type WishlistRequest struct {
	ProductId uint64 `json:"product_id" validate:"required"`
}
