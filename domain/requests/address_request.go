package requests

type AddressRequest struct {
	Country      string `validate:"required,alpha" json:"country,omitempty"`
	City         string `validate:"required,alpha" json:"city,omitempty"`
	ReceiverName string `validate:"required,alpha" json:"receiver_name,omitempty"`
	PhoneNumber  string `validate:"required,numeric,max=20" json:"phone_number,omitempty"`
	IsPrimary    bool   `validate:"required" json:"is_primary,omitempty"`
}
