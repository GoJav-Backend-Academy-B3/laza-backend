package midtrans_core

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/consts"
)

var MidtransCore coreapi.Client

func init() {
	MidtransCore.ServerKey = consts.MidtransSandBoxServerKey
	MidtransCore.ClientKey = consts.MidtransSandBoxClientKey
	MidtransCore.Env = midtrans.Sandbox
}
