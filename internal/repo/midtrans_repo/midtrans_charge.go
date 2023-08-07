package midtrans_repo

import (
	"errors"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/external/midtrans"
)

func ChargeMidtrans(req *coreapi.ChargeReq) (coreapi.ChargeResponse, error) {
	chargeResp, errmd := midtrans_core.MidtransCore.ChargeTransaction(req)

	if errmd != nil {
		return *chargeResp, errors.New(errmd.Message)
	}

	return *chargeResp, errors.New(errmd.Message)
}
