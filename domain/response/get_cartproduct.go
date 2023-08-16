package response

type CartPorduct struct {
	Id        uint64  `json:"id"`
	Name      string  `json:"product_name"`
	ImageUrl  string  `json:"image_url"`
	Price     float64 `json:"price"`
	BrandName string  `json:"brand_name"`
	Quantity  int     `json:"quantity"`
	Size      string  `json:"size"`
}

type CartOrderInfo struct {
	SubTotal     float64 `json:"sub_total"`
	ShippingCost float64 `json:"shipping_cost"`
	Total        float64 `json:"total"`
}

type CartInfo struct {
	CartPorduct   []CartPorduct `json:"products"`
	CartOrderInfo CartOrderInfo `json:"order_info"`
}
