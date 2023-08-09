package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type InsertBanksUsecase struct {
	insertBankAction repositories.InsertAction[model.Bank]
	bankExistsAction action.ExistsBank
}

func NewInsertBankUsecase(repoBank repositories.InsertAction[model.Bank], repoExistsBank action.ExistsBank) bank.InsertBankUsecase {
	return &InsertBanksUsecase{
		insertBankAction: repoBank,
		bankExistsAction: repoExistsBank,
	}
}

func (uc *InsertBanksUsecase) Execute(request requests.BankRequest) *helper.Response {
	if bankExists := uc.bankExistsAction.ExistsBank(request.BankName); bankExists {
		return helper.GetResponse("Name Bank is already registered", 401, true)
	}

	file, err := request.LogoUrl.Open()
	if err != nil {
		// TODO: Should return error here
		return helper.GetResponse(err.Error(), 500, true)
	}
	defer file.Close()

	url, err := helper.UploadImageFile("bank", file)
	if err != nil {
		// TODO: Should return error here
		return helper.GetResponse(err.Error(), 500, true)
	}

	bank := model.Bank{
		BankName: request.BankName,
		BankCode: request.BankCode,
		LogoUrl:  url,
	}

	result, err := uc.insertBankAction.Insert(bank)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse(result, 201, true)
}
