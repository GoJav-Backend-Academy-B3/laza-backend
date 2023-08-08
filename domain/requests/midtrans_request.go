package requests

type ChargeGopay struct {
	PaymentType        string            `json:"payment_type"`
	TransactionDetails TransactionDetail `json:"transaction_details"`
	Gopay              Gopay             `json:"gopay"`
}

type TransactionDetail struct {
	GrossAmount int    `json:"gross_amount"`
	OrderId     string `json:"order_id"`
}

type Gopay struct {
	EnableCallback bool   `json:"enable_callback"`
	CallbackUrl    string `json:"callback_url"`
}
