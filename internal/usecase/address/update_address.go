package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type updateAddressUsecase struct {
	updateAddress repositories.UpdateAction[model.Address]
	getById       repositories.GetByIdAction[model.Address]
}

// UpdateAddressById implements address.UpdateAddressInterface.
func (u *updateAddressUsecase) UpdateAddressById(id uint64, userId uint64, request requests.AddressRequest) (address model.Address, err error) {
	address, err = u.getById.GetById(id)
	if err != nil {
		return
	}

	address = model.Address{
		Country:      request.Country,
		City:         request.City,
		ReceiverName: request.ReceiverName,
		PhoneNumber:  request.PhoneNumber,
		IsPrimary:    request.IsPrimary,
		UserId:       userId,
	}

	address, err = u.updateAddress.Update(id, address)
	if err != nil {
		return
	}

	return
}

func NewUpdateAddressUsecase(updateAddress repositories.UpdateAction[model.Address], getById repositories.GetByIdAction[model.Address]) address.UpdateAddressUsecase {
	return &updateAddressUsecase{
		updateAddress: updateAddress,
		getById:       getById,
	}
}
