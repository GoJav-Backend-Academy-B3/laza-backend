package response

import "github.com/phincon-backend/laza/domain/model"

type TransactionBankOrderResponse struct {
	BasicOrderResponse
	TransactionBankId uint64 `json:"va_number"`
}

func (r *TransactionBankOrderResponse) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.OrderStatus = m.OrderStatus
	r.AddressId = m.AddressId
	//if m.TransactionBankId.Valid {
	//	r.TransactionBankId = uint64(m.TransactionBankId.Int64)
	//} else {
	//	panic("TransactionBankId is nil! Cannot fill.")
	//}
}
