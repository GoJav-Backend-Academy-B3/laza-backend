package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/address"
	"github.com/phincon-backend/laza/domain/usecases/address"
	repository "github.com/phincon-backend/laza/internal/repo/address"
)

type getAddressUsecase struct {
	getAllAddress  action.GetAllByUserIdAction
	getAddressById repositories.GetByIdAction[model.Address]
}

// GetAddressById implements address.GetAddressUsecase.
func (u *getAddressUsecase) GetAddressById(id uint64) (address model.Address, err error) {
	address, err = u.getAddressById.GetById(id)
	if err != nil {
		return
	}

	return
}

// GetAllAddressByUserId implements address.GetAddressUsecase.
func (u *getAddressUsecase) GetAllAddressByUserId(userId uint64) (addresses []model.Address, err error) {
	addresses, err = u.getAllAddress.GetAllByUserId(userId)
	if err != nil {
		return
	}

	return
}

func NewGetAddrressUsecase(addressRepo repository.AddressRepo) address.GetAddressUsecase {
	return &getAddressUsecase{getAllAddress: &addressRepo, getAddressById: &addressRepo}
}
