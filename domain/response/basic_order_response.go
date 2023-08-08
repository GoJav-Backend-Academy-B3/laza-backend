package response

import (
	"time"
)

type BasicOrderResponse struct {
	Id            string
	Amount        int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserId        uint64
	OrderStatusId uint64
	AddressId     uint64
}
