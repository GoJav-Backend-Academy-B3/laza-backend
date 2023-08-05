package response

type CartPorduct struct {
	Id        uint64  `json:"id"`
	Name      string  `json:"product_name"`
	ImageUrl  string  `json:"image_url"`
	Price     float64 `json:"price"`
	BrandName string  `json:"brand_name"`
	Quantity  int     `json:"quantity"`
}
