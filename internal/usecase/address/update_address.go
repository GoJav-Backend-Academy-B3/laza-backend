package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/address"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/address"
	repository "github.com/phincon-backend/laza/internal/repo/address"
)

type updateAddressUsecase struct {
	updateAddress                          repositories.UpdateAction[model.Address]
	getById                                repositories.GetByIdAction[model.Address]
	setAllprimaryAddressesNonPrimaryAction action.SetAllAddressesNonPrimaryAction
}

// UpdateAddressById implements address.UpdateAddressInterface.
func (u *updateAddressUsecase) UpdateAddressById(id uint64, userId uint64, request requests.AddressRequest) (address model.Address, err error) {
	address, err = u.getById.GetById(id)
	if err != nil {
		return
	}

	address = model.Address{
		Id:           id,
		Country:      request.Country,
		City:         request.City,
		ReceiverName: request.ReceiverName,
		PhoneNumber:  request.PhoneNumber,
		IsPrimary:    request.IsPrimary,
		UserId:       userId,
	}

	if address.IsPrimary {
		err := u.setAllprimaryAddressesNonPrimaryAction.SetAllAddressesNonPrimary(address.UserId)
		if err != nil {
			return address, err
		}
	}

	address, err = u.updateAddress.Update(id, address)
	if err != nil {
		return
	}

	return
}

func NewUpdateAddressUsecase(addressRepo repository.AddressRepo) address.UpdateAddressUsecase {
	return &updateAddressUsecase{
		updateAddress:                          &addressRepo,
		getById:                                &addressRepo,
		setAllprimaryAddressesNonPrimaryAction: &addressRepo,
	}
}
