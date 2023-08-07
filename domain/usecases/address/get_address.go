package address

import "github.com/phincon-backend/laza/domain/model"

type GetAddressUsecase interface {
	GetAllAddressByUserId(userId uint64) (addresses []model.Address, err error)
	GetAddressById(id uint64) (address model.Address, err error)
}
