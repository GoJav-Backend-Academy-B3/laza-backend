package response

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
)

type GopayPaymentResponse struct {
	Id         uint64    `json:"id,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	DeepLink   string    `json:"deep_link,omitempty"`
	QRCode     string    `json:"qr_code,omitempty"`
	ExpiryTime time.Time `json:"expiry_time"`
}

func (r *GopayPaymentResponse) FillFromEntity(m *model.Gopay) {
	r.Id = m.Id
	r.Phone = m.Phone
	r.DeepLink = m.Deeplink
	r.QRCode = m.QRCode
	r.ExpiryTime = m.ExpiryTime
}
