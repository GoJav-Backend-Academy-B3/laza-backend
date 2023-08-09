package requests

type OrderBasic struct {
	AddressId int            `json:"address_id" binding:"required"`
	Products  []ProductOrder `json:"products" binding:"required"`
}

type OrderWithGopay struct {
	OrderBasic
	CallbackUrl string `json:"callback_url" binding:"required"`
}

type OrderWithBank struct {
	OrderBasic
	Bank string `json:"bank" binding:"required"`
}

type OrderWithCC struct {
	OrderBasic
	CreditCard CreditCardOrder `json:"credit_card" binding:"required"`
}

type CreditCardOrder struct {
	Id         int    `json:"id" binding:"required"`
	CardNumber string `json:"card_number" binding:"required"`
	ExpMonth   int    `json:"exp_month" binding:"required"`
	ExpYear    int    `json:"exp_year" binding:"required"`
	CVV        string `json:"cvv" binding:"required"`
}
