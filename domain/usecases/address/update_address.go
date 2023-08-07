package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/requests"
)

type UpdateAddressUsecase interface {
	UpdateAddressById(id uint64, userId uint64, request requests.AddressRequest) (address model.Address, err error)
}
