package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type UpdateBankUsecase struct {
	updateAction action.UpdateBank[model.Bank]
}

func NewUpdateBankUsecase(repo action.UpdateBank[model.Bank]) bank.UpdateBankUsecase {
	return &UpdateBankUsecase{updateAction: repo}
}

func (uc *UpdateBankUsecase) Execute(id uint64, bank model.Bank) *helper.Response {
	result, err := uc.updateAction.Update(id, bank)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
