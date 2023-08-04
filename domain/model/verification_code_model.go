package model

import "time"

type VerificationCode struct {
	Id         uint      `json:"id,omitempty" gorm:"primarykey"`
	Code       string    `json:"code,omitempty"`
	ExpiryDate time.Time `json:"expiry_date"`
	UserId     uint64    `json:"user_id,omitempty"`
}
