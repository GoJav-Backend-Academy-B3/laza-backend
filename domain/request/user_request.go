package request

type User struct {
	FullName string `json:"full_name" form:"full_name" validate:"required,alpha"`
	Username string `json:"username" form:"username" validate:"required,alphanum"`
	Password string `json:"password" form:"password" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Image    string `json:"image" form:"image"`
}

type Login struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

type Resend struct {
	Email string `json:"email" validate:"required,email"`
}
