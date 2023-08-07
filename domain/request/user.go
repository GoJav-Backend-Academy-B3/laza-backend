package request

type User struct {
	FullName string `json:"full_name" form:"full_name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Image    string `json:"image" form:"image"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Resend struct {
	Email string `json:"email"`
}
