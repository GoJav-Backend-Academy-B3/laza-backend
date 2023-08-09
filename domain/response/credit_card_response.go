package response

import "github.com/phincon-backend/laza/domain/model"

type CreditCardResponse struct {
	Id           uint64 `json:"id,omitempty"  gorm:"primarykey"`
	CardNumber   string `json:"card_number,omitempty"`
	ExpiredMonth int    `json:"expired_month,omitempty"`
	ExpiredYear  int    `json:"expired_year,omitempty"`
	UserId       uint64 `json:"user_id,omitempty"`
}

func (cc CreditCardResponse) FillFromEntity(md model.CreditCard) CreditCardResponse {
	return CreditCardResponse{
		Id:           md.Id,
		CardNumber:   md.CardNumber,
		ExpiredMonth: md.ExpiredMonth,
		ExpiredYear:  md.ExpiredYear,
		UserId:       md.UserId,
	}
}
