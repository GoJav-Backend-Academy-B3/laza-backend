package model

import "time"

type User struct {
	Id         uint      `json:"id,omitempty"`
	Username   string    `json:"username,omitempty"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email,omitempty"`
	FullName   string    `json:"full_name,omitempty"`
	IsVerified bool      `json:"is_verified,omitempty"`
	IsAdmin    bool      `json:"is_admin,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
