package address

import (
	"strconv"

	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/address"
)

type deleteAddressUsecase struct {
	delete repositories.DeleteAction[model.Address]
}

// DeleteAddressById implements address.DeleteAddressInterface.
func (u *deleteAddressUsecase) DeleteAddressById(id uint64) (err error) {
	convert := strconv.Itoa(int(id))

	err = u.delete.Delete(convert)
	if err != nil {
		return
	}

	return
}

func NewDeleteAddressUsecase(delete repositories.DeleteAction[model.Address]) address.DeleteAddressUsecase {
	return &deleteAddressUsecase{delete: delete}
}
