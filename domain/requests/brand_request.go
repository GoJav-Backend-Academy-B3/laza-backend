package requests

import "mime/multipart"

type BrandRequest struct {
	Name    string                `json:"name,omitempty" validate:"required"`
	LogoUrl *multipart.FileHeader `json:"logo_url,omitempty" swaggerignore:"true"`
}
