package response

import "github.com/phincon-backend/laza/domain/model"

type CreditCardOrderResponse struct {
	BasicOrderResponse
	CreditCardId uint64
}

func (r *CreditCardOrderResponse) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.OrderStatusId = m.OrderStatusId
	r.AddressId = m.AddressId
	if m.CreditCardId.Valid {
		r.CreditCardId = uint64(m.CreditCardId.Int64)
	} else {
		panic("CreditCardId is nil! Cannot fill.")
	}
}
