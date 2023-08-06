package address

import (
	"strconv"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type updateAddressUsecase struct {
	updateAddress repositories.UpdateAction[model.Address]
}

// UpdateAddressById implements address.UpdateAddressInterface.
func (u *updateAddressUsecase) UpdateAddressById(id uint64, request requests.AddressRequest) (address model.Address, err error) {
	convert := strconv.Itoa(int(id))

	address = model.Address{
		Country:      request.Country,
		City:         request.City,
		ReceiverName: request.ReceiverName,
		PhoneNumber:  request.PhoneNumber,
		IsPrimary:    request.IsPrimary,
		UserId:       request.UserId,
	}

	address, err = u.updateAddress.Update(convert, address)
	if err != nil {
		return
	}

	return
}

func NewUpdateAddressUsecase(updateAddress repositories.UpdateAction[model.Address]) address.UpdateAddressUsecase {
	return &updateAddressUsecase{updateAddress: updateAddress}
}
