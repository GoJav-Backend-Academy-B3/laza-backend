package requests

type AddressRequest struct {
	Country      string `validate:"required" json:"country,omitempty"`
	City         string `validate:"required" json:"city,omitempty"`
	ReceiverName string `validate:"required" json:"receiver_name,omitempty"`
	PhoneNumber  string `validate:"required,numeric,max=20" json:"phone_number,omitempty"`
	IsPrimary    bool   `validate:"required" json:"is_primary,omitempty"`
}
