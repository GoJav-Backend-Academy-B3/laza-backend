package midtrans

import (
	"github.com/midtrans/midtrans-go/coreapi"
)

type ChargeMidtransAction interface {
	ChargeMidtrans(req *coreapi.ChargeReq) (coreapi.ChargeResponse, error)
	//ChargeGopay(req requests.ChargeGopay) (coreapi.ChargeResponse, error)
}
