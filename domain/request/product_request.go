package request

import "mime/multipart"

type ProductRequest struct {
	Name        string                `binding:"required,printascii" form:"name"`
	Description string                `binding:"required,printascii" form:"description"`
	Image       *multipart.FileHeader `binding:"required" form:"image"`
	Price       float64               `binding:"required,numeric" form:"price"`
	Category    string                `binding:"required,printascii" form:"category"`
	Brand       string                `binding:"required,printascii" form:"brand"`
	Sizes       []string              `binding:"required,printascii" form:"sizes"`
}
