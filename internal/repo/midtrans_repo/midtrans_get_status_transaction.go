package midtrans_repo

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	midtrans_core "github.com/phincon-backend/laza/external/midtrans"
)

func FetchMidtransTransactionStatus(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {

	transactionStatus, err := midtrans_core.MidtransCore.CheckTransaction(orderId)

	if err != nil {
		return nil, err
	}
	return transactionStatus, err
}
