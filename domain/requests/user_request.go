package requests

type User struct {
	FullName string `json:"full_name" form:"full_name" validate:"required,min=3"`
	Username string `json:"username" form:"username" validate:"required,min=3"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Image    string `json:"image" form:"image"`
}

type Login struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type Email struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdatePassword struct {
	Password   string `json:"password" validate:"required,min=8"`
	RePassword string `json:"re_password" validate:"required,min=8"`
}
