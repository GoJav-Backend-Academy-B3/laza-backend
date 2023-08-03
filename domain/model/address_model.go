package model

type Address struct {
	Id           uint64 `json:"id,omitempty"`
	Country      string `json:"country,omitempty"`
	City         string `json:"city,omitempty"`
	RecieverName string `json:"reciever-name,omitempty"`
	PhoneNumber  string `json:"phone-number,omitempty"`
	IsPrimary    bool   `json:"is-primary,omitempty"`
	UserId       uint   `json:"user-id,omitempty"`
}
