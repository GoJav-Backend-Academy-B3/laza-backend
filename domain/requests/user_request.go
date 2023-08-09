package requests

import "mime/multipart"

type Register struct {
	FullName string                `json:"full_name" form:"full_name" validate:"required,min=3"`
	Username string                `json:"username" form:"username" validate:"required,alphanum,min=3"`
	Password string                `json:"password" form:"password" validate:"required,min=8"`
	Email    string                `json:"email" form:"email" validate:"required,email"`
	Image    *multipart.FileHeader `json:"image" form:"image" swaggerignore:"true"`
}

type UpdateUser struct {
	FullName string                `json:"full_name" form:"full_name" validate:"required,min=3"`
	Username string                `json:"username" form:"username" validate:"required,alphanum,min=3"`
	Email    string                `json:"email" form:"email" validate:"required,email"`
	Image    *multipart.FileHeader `json:"image" form:"image" swaggerignore:"true"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
	RePassword  string `json:"re_password" validate:"required,min=8"`
}

type Login struct {
	Username string `json:"username" validate:"required,alphanum,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type VerificationCode struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,min=4"`
}

type ResetPassword struct {
	NewPassword string `json:"new_password" validate:"required,min=8"`
	RePassword  string `json:"re_password" validate:"required,min=8"`
}

type Email struct {
	Email string `json:"email" validate:"required,email"`
}
