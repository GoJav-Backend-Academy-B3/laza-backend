package requests

import "github.com/phincon-backend/laza/domain/model"

type CreditCardRequest struct {
	CardNumber   string `json:"card_number" validate:"required,number,min=16,max=16"`
	ExpiredMonth int    `json:"expired_month" validate:"required"`
	ExpiredYear  int    `json:"expired_year" validate:"required"`
	CVV          string `json:"cvv"`
}

func (cc CreditCardRequest) FilltoField(userId uint64, dt *model.CreditCard) {
	dt.CardNumber = cc.CardNumber
	dt.ExpiredYear = cc.ExpiredYear
	dt.ExpiredMonth = cc.ExpiredMonth
	dt.UserId = userId
}

func RequestFiledToCardOrder(userId uint64, rq CreditCardOrder, dt *model.CreditCard) {
	dt.CardNumber = rq.CardNumber
	dt.ExpiredYear = rq.ExpYear
	dt.ExpiredMonth = rq.ExpMonth
	dt.UserId = userId
}
