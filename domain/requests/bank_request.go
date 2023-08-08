package requests

import "mime/multipart"

type BankRequest struct {
	BankName string                `json:"bank_name,omitempty" binding:"required" form:"bank_name" `
	BankCode string                `json:"bank_code,omitempty" binding:"required,numeric" form:"bank_code" `
	LogoUrl  *multipart.FileHeader `json:"image,omitempty" form:"image" swaggerignore:"true"`
}
