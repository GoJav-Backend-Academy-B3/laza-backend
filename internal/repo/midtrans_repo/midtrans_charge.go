package midtrans

import (
	"errors"
	"fmt"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (m *MidtransRepo) ChargeMidtrans(req *coreapi.ChargeReq) (coreapi.ChargeResponse, error) {
	fmt.Println("repo request: ", req)
	fmt.Println("midtranscore: ", *m)
	fmt.Println("midclient: ", *m.midtransClient)
	fmt.Println("midclient charge: ", m.midtransClient.ChargeTransaction)
	chargeResp, errmd := m.midtransClient.ChargeTransaction(req)
	fmt.Println("charge respond: ", chargeResp)
	if errmd != nil {
		return *chargeResp, errors.New(errmd.Message)
	}

	return *chargeResp, nil
}
