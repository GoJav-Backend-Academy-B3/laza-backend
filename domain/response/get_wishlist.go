package response

type WishlistProduct struct {
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}
