package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type FetchMidtransTransactionAction interface {
	FetchMidtransTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error)
}
