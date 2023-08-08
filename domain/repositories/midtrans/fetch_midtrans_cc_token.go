package midtrans

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type FetchMidtransCCTokenAction interface {
	FetchMidtransCCToken(cardNumber string, expMonth int, expYear int, cvv string) (*coreapi.CardTokenResponse, *midtrans.Error)
}
