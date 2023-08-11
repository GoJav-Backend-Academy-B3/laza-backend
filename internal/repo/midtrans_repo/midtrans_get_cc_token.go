package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/phincon-backend/laza/config"
)

func (m *MidtransRepo) FetchMidtransCCToken(cardNumber string, expMonth int, expYear int, cvv string) (*coreapi.CardTokenResponse, *midtrans.Error) {

	token, err := m.midtransClient.CardToken(cardNumber, expMonth, expYear, cvv, config.MidtransCore.ClientKey)
	if err != nil {
		return nil, err
	}
	return token, err
}
