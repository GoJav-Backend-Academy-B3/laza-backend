package midtrans

import "github.com/midtrans/midtrans-go/coreapi"

type MidtransRepo struct {
	midtransClient *coreapi.Client
}

func NewMidtransRepo(midtransClient *coreapi.Client) *MidtransRepo {
	return &MidtransRepo{midtransClient: midtransClient}
}
