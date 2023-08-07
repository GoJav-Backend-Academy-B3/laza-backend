package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type InsertBanksUsecase struct {
	insertBankAction action.InsertBank[model.Bank]
	bankExistsAction action.ExistsBank
}

func NewInsertBankUsecase(repoBank action.InsertBank[model.Bank], repoExistsBank action.ExistsBank) bank.InsertBankUsecase {
	return &InsertBanksUsecase{
		insertBankAction: repoBank,
		bankExistsAction: repoExistsBank,
	}
}

func (uc *InsertBanksUsecase) Execute(bank model.Bank) *helper.Response {
	if bank.BankName == "" {
		return helper.GetResponse("Name Bank is required", 400, true)
	}
	if bank.BankCode == "" {
		return helper.GetResponse("Bank Code is required", 400, true)
	}
	if userExists := uc.bankExistsAction.ExistsBank(bank.BankName); userExists {
		return helper.GetResponse("Name Bank is already registered", 401, true)
	}
	result, err := uc.insertBankAction.Insert(bank)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 200, true)
}
