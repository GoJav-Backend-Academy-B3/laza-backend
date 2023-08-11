package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/address"
	usecase "github.com/phincon-backend/laza/domain/usecases/address"
	repository "github.com/phincon-backend/laza/internal/repo/address"
)

type deleteAddressUsecase struct {
	delete            repositories.DeleteAction[model.Address]
	getById           repositories.GetByIdAction[model.Address]
	findLatestAddress action.FindLatestAddressByUserIdAction
	update            repositories.UpdateAction[model.Address]
}

// DeleteAddressById implements address.DeleteAddressInterface.
func (u *deleteAddressUsecase) DeleteAddressById(id uint64) (err error) {
	address, err := u.getById.GetById(id)
	if err != nil {
		return
	}

	if address.IsPrimary {
		err = u.delete.Delete(id)
		if err != nil {
			return
		}

		result := u.findLatestAddress.FindLatestAddressByUserId(address.UserId)
		if err != nil {
			return
		}

		result.IsPrimary = true

		_, err := u.update.Update(result.Id, result)
		if err != nil {
			return err
		}

	} else {
		err = u.delete.Delete(id)
		if err != nil {
			return
		}
	}

	return
}

func NewDeleteAddressUsecase(addressRepo repository.AddressRepo) usecase.DeleteAddressUsecase {
	return &deleteAddressUsecase{
		delete:            &addressRepo,
		getById:           &addressRepo,
		findLatestAddress: &addressRepo,
		update:            &addressRepo,
	}
}
