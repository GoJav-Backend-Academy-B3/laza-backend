package midtrans

//import (
//	"errors"
//	"github.com/midtrans/midtrans-go/coreapi"
//	"github.com/phincon-backend/laza/domain/requests"
//	midtrans_core "github.com/phincon-backend/laza/external/midtrans"
//	"net/http"
//)
//
//func (m *MidtransRepo) ChargeGopay(req requests.ChargeGopay) (coreapi.ChargeResponse, error) {
//	chargeResp, errmd := midtrans_core.MidtransCore.ChargeTransaction(req)
//
//
//
//	respond, err := http.Post(
//		"https://api.sandbox.midtrans.com/v2/charge",
//		"application/json",
//		io.b,
//		,)
//
//	return *chargeResp, errors.New(errmd.Message)
//}
