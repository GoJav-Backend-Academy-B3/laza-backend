package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	repository "github.com/phincon-backend/laza/domain/repositories/address"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type addAddressUsecase struct {
	addressRepo                            repositories.InsertAction[model.Address]
	setAllprimaryAddressesNonPrimaryAction repository.SetAllAddressesNonPrimaryAction
}

func NewAddAddressUsecase(addressRepo repositories.InsertAction[model.Address], setAllprimaryAddressesNonPrimaryAction repository.SetAllAddressesNonPrimaryAction) address.AddAddressUsecase {
	return &addAddressUsecase{
		addressRepo:                            addressRepo,
		setAllprimaryAddressesNonPrimaryAction: setAllprimaryAddressesNonPrimaryAction,
	}
}

// AddAddress implements address.AddAddressUsecase.
func (u *addAddressUsecase) AddAddress(request requests.AddressRequest) (model.Address, error) {
	address := model.Address{
		Country:      request.Country,
		City:         request.City,
		ReceiverName: request.ReceiverName,
		PhoneNumber:  request.PhoneNumber,
		IsPrimary:    request.IsPrimary,
		UserId:       request.UserId,
	}

	if address.IsPrimary {
		err := u.setAllprimaryAddressesNonPrimaryAction.SetAllAddressesNonPrimary(address.UserId)
		if err != nil {
			return address, err
		}
	}

	newAddress, err := u.addressRepo.Insert(address)
	if err != nil {
		return newAddress, err
	}

	return newAddress, nil
}
