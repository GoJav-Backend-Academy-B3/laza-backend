package midtrans

import (
	"errors"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (m *MidtransRepo) ChargeMidtrans(req *coreapi.ChargeReq) (coreapi.ChargeResponse, error) {
	chargeResp, errmd := m.midtransClient.ChargeTransaction(req)

	if errmd != nil {
		return *chargeResp, errors.New(errmd.Message)
	}

	return *chargeResp, nil
}
