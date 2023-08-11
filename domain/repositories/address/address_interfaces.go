package address

import "github.com/phincon-backend/laza/domain/model"

type GetAllByUserIdAction interface {
	GetAllByUserId(userId uint64) (addresses []model.Address, err error)
}

type SetAllAddressesNonPrimaryAction interface {
	SetAllAddressesNonPrimary(userId uint64) (err error)
}

type FindLatestAddressByUserIdAction interface {
	FindLatestAddressByUserId(user_id uint64) (address model.Address)
}
