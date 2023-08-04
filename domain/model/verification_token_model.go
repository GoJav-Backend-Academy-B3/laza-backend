package model

import "time"

type VerificationToken struct {
	Id         uint      `json:"id,omitempty" gorm:"primarykey"`
	Token      string    `json:"token,omitempty"`
	ExpiryDate time.Time `json:"expiry_date"`
	UserId     uint64    `json:"user_id,omitempty"`
}

func (VerificationToken) TableName() string {
	return "verification_token"
}
