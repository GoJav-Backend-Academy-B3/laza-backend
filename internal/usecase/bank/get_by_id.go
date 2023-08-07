package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
	"gorm.io/gorm"
)

type GetByIdBankUsecase struct {
	getBankByIdAction action.GetBankByIdAction[model.Bank]
}

func NewGetByIdBankUsecase(repo action.GetBankByIdAction[model.Bank]) bank.GetByIdBankUsecase {
	return &GetByIdBankUsecase{getBankByIdAction: repo}
}

func (uc *GetByIdBankUsecase) Execute(id uint64) *helper.Response {
	result, err := uc.getBankByIdAction.GetBankById(id)
	if err != nil || err == gorm.ErrRecordNotFound {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
