package config

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"os"
)

var MidtransCore *coreapi.Client

func Init() *coreapi.Client {
	c := coreapi.Client{}
	MidtransCore = &c

	clientKey := os.Getenv("MIDTRANS_CLIENT_KEY")
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")

	MidtransCore.New(serverKey, midtrans.Sandbox)
	MidtransCore.ServerKey = serverKey
	MidtransCore.ClientKey = clientKey
	MidtransCore.Env = midtrans.Sandbox
	MidtransCore.Options = &midtrans.ConfigOptions{}
	return MidtransCore
}
