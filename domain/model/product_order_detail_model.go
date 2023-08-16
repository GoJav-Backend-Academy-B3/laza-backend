package model

type ProductOrderDetail struct {
	Id          uint64 `json:"id,omitempty" gorm:"primarykey"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	Price       int    `json:"price,omitempty"`
	Category    string `json:"category_name,omitempty"`
	BrandName   string `json:"brand_name,omitempty"`
	Quantity    int    `json:"quantity"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	OrderId     string `json:"order_id,omitempty"`
}

func (ProductOrderDetail) TableName() string {
	return "product_order_detail"
}
