package response

import (
	"time"

	"github.com/phincon-backend/laza/domain/model"
)

type GopayPaymentResponse struct {
	Id            uint64
	Phone         string
	DeepLink      string
	QRCode        string
	GetStatusLink string
	CancelLink    string
	ExpiryTime    time.Time
}

func (r *GopayPaymentResponse) FillFromEntity(m *model.Gopay) {
	r.Id = m.Id
	r.Phone = m.Phone
	r.DeepLink = m.Deeplink
	r.QRCode = m.QRCode
	r.GetStatusLink = m.GetStatusLink
	r.CancelLink = m.CancelLink
	r.ExpiryTime = m.ExpiryTime
}
