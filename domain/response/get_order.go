package response

import (
	"github.com/phincon-backend/laza/domain/model"
	"time"
)

type Order struct {
	Id                string    `json:"id"`
	Amount            int64     `json:"amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	UserId            uint64    `json:"user_id"`
	OrderStatus       string    `json:"order_status"`
	AddressId         uint64    `json:"address_id"`
	CreditCardId      int64     `json:"credit_card_id,omitempty"`
	TransactionBankId int64     `json:"transaction_bank_id,omitempty"`
	GopayId           int64     `json:"gopay_id,omitempty"`
}

func (r *Order) FillFromEntity(m *model.Order) {
	r.Id = m.Id
	r.Amount = m.Amount
	r.CreatedAt = m.CreatedAt
	r.UpdatedAt = m.UpdatedAt
	r.AddressId = m.AddressId
	r.UserId = m.UserId
	r.OrderStatus = m.OrderStatus
	//r.GopayId = m.GopayId.Int64
	//r.CreditCardId = m.CreditCardId.Int64
	//r.TransactionBankId = m.TransactionBankId.Int64
}
