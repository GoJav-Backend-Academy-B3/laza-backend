package response

type WishlistProduct struct {
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}

type WishProduct struct {
	Id        uint64  `json:"id"`
	Name      string  `json:"name"`
	ImageUrl  string  `json:"image_url"`
	Price     float64 `json:"price"`
	BrandName string  `json:"brand_name"`
}

type WishListProductLimit struct {
	Total   int               `json:"total"`
	Product []ProductOverview `json:"products"`
}
