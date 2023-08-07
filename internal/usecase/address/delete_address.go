package address

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type deleteAddressUsecase struct {
	delete  repositories.DeleteAction[model.Address]
	getById repositories.GetByIdAction[model.Address]
}

// DeleteAddressById implements address.DeleteAddressInterface.
func (u *deleteAddressUsecase) DeleteAddressById(id uint64) (err error) {
	_, err = u.getById.GetById(id)
	if err != nil {
		return
	}

	err = u.delete.Delete(id)
	if err != nil {
		return
	}

	return
}

func NewDeleteAddressUsecase(delete repositories.DeleteAction[model.Address], getById repositories.GetByIdAction[model.Address]) address.DeleteAddressUsecase {
	return &deleteAddressUsecase{
		delete:  delete,
		getById: getById,
	}
}
