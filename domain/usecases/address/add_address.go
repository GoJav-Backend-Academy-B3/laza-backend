package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type AddAddressUsecase interface {
	AddAddress(request requests.AddressRequest) (model.Address, error)
}
