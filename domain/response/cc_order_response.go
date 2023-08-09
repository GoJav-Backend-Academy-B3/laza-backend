package response

import "github.com/phincon-backend/laza/domain/model"

type CreditCardOrderResponse struct {
	BasicOrderResponse
	CreditCardId uint64 `json:"credit_card_id"`
}

func (r *CreditCardOrderResponse) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.OrderStatus = m.OrderStatus
	r.AddressId = m.AddressId
	if m.CreditCardId.Valid {
		r.CreditCardId = uint64(m.CreditCardId.Int64)
	} else {
		panic("CreditCardId is nil! Cannot fill.")
	}
}
