package response

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
)

type User struct {
	Id         uint      `json:"id"`
	FullName   string    `json:"full_name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	ImageUrl   string    `json:"image_url"`
	IsVerified bool      `json:"is_verified,omitempty"`
	IsAdmin    bool      `json:"is_admin,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func UserModelResponse(data model.User) *User {
	return &User{
		Id:         data.Id,
		FullName:   data.FullName,
		Username:   data.Username,
		Email:      data.Email,
		ImageUrl:   data.ImageUrl,
		IsVerified: data.IsVerified,
		IsAdmin:    data.IsAdmin,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}
