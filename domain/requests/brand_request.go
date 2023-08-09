package requests

import "mime/multipart"

type BrandRequest struct {
	Name    string                `json:"name,omitempty" binding:"required" form:"name"`
	LogoUrl *multipart.FileHeader `json:"logo_url,omitempty" swaggerignore:"true" form:"logo_url"`
}
