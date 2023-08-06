package response

import "time"

type User struct {
	Id         uint      `json:"id"`
	FullName   string    `json:"full_name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	ImageUrl   string    `json:"image_url"`
	IsVerified bool      `json:"is_verified,omitempty"`
	IsAdmin    bool      `json:"is_admin,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}