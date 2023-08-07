package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/external/midtrans"
)

func (m *MidtransRepo) FetchMidtransTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {

	transactionStatus, err := midtrans_core.MidtransCore.CheckTransaction(orderId)

	if err != nil {
		return nil, err
	}
	return transactionStatus, err
}
