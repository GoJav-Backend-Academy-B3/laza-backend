package request

type OrderWithGopay struct {
	AddressId   int            `json:"address_id" validate:"required"`
	Products    []ProductOrder `json:"products" validate:"required"`
	CallbackUrl string         `json:"callback_url" validate:"required"`
}

type OrderWithBank struct {
	AddressId int            `json:"address_id" validate:"required"`
	Products  []ProductOrder `json:"products" validate:"required"`
	Bank      string         `json:"bank" validate:"required"`
}
