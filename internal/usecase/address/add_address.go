package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/address"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/address"
	repository "github.com/phincon-backend/laza/internal/repo/address"
)

type addAddressUsecase struct {
	addressRepo                            repositories.InsertAction[model.Address]
	setAllprimaryAddressesNonPrimaryAction action.SetAllAddressesNonPrimaryAction
}

func NewAddAddressUsecase(addressRepo repository.AddressRepo) address.AddAddressUsecase {
	return &addAddressUsecase{
		addressRepo:                            &addressRepo,
		setAllprimaryAddressesNonPrimaryAction: &addressRepo,
	}
}

// AddAddress implements address.AddAddressUsecase.
func (u *addAddressUsecase) AddAddress(request requests.AddressRequest, userId uint64) (model.Address, error) {
	address := model.Address{
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

	newAddress, err := u.addressRepo.Insert(address)
	if err != nil {
		return newAddress, err
	}

	return newAddress, nil
}
