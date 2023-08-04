package model

type Address struct {
	Id           uint64  `json:"id,omitempty" gorm:"primarykey"`
	Country      string  `json:"country,omitempty"`
	City         string  `json:"city,omitempty"`
	RecieverName string  `json:"reciever_name,omitempty"`
	PhoneNumber  string  `json:"phone_number,omitempty"`
	IsPrimary    bool    `json:"is_primary,omitempty"`
	UserId       uint64  `json:"user_id,omitempty"`
	Orders       []Order `json:"orders,omitempty" gorm:"foreignkey:Id"`
}

func (Address) TableName() string {
	return "address"
}
