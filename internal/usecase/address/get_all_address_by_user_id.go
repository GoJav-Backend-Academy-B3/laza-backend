package address

import (
	"strconv"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	repository "github.com/phincon-backend/laza/domain/repositories/address"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type getAddressUsecase struct {
	getAllAddress  repository.GetAllByUserIdAction
	getAddressById repositories.GetByIdAction[model.Address]
}

// GetAddressById implements address.GetAddressUsecase.
func (u *getAddressUsecase) GetAddressById(id uint64) (address model.Address, err error) {
	convert := strconv.Itoa(int(id))
	address, err = u.getAddressById.GetById(convert)
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

func NewGetAddrressUsecase(getAllAddress repository.GetAllByUserIdAction, getAddressById repositories.GetByIdAction[model.Address]) address.GetAddressUsecase {
	return &getAddressUsecase{getAllAddress: getAllAddress, getAddressById: getAddressById}
}
