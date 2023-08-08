package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func (m *MidtransRepo) FetchMidtransTransaction(orderId string) (*coreapi.TransactionStatusResponse, *midtrans.Error) {

	transactionStatus, err := m.midtransClient.CheckTransaction(orderId)

	if err != nil {
		return nil, err
	}
	return transactionStatus, err
}
