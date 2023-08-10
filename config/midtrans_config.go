package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/consts"
)

var MidtransCore *coreapi.Client

func Init() *coreapi.Client {
	c := coreapi.Client{}
	MidtransCore = &c
	MidtransCore.New(consts.MidtransSandBoxServerKey, midtrans.Sandbox)
	MidtransCore.New(consts.MidtransSandBoxServerKey, midtrans.Sandbox)
	MidtransCore.ServerKey = consts.MidtransSandBoxServerKey
	MidtransCore.ClientKey = consts.MidtransSandBoxClientKey
	MidtransCore.Env = midtrans.Sandbox
	MidtransCore.Options = &midtrans.ConfigOptions{}
	return MidtransCore
}
