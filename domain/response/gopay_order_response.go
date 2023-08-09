package response

import "github.com/phincon-backend/laza/domain/model"

type GopayOrderResponse struct {
	BasicOrderResponse
	GopayId uint64 `json:"gopay_id"`
}

func (r *GopayOrderResponse) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.OrderStatus = m.OrderStatus
	r.AddressId = m.AddressId
	if m.GopayId.Valid {
		r.GopayId = uint64(m.GopayId.Int64)
	} else {
		panic("GopayId is nil! Cannot fill.")
	}
}
