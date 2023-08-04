package model

type OrderStatus struct {
	Id     uint64  `json:"id,omitempty" gorm:"primarykey"`
	Status string  `json:"status,omitempty"`
	Orders []Order `json:"orders" gorm:"foreignkey:Id"`
}
