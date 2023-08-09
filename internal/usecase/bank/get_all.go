package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	"github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type GetAllBankUsecase struct {
	getAllBank repositories.GetAllAction[model.Bank]
}

func NewGetAllBankUsecase(repo repositories.GetAllAction[model.Bank]) bank.GetAllBankUsecase {
	return &GetAllBankUsecase{getAllBank: repo}
}

func (uc *GetAllBankUsecase) Execute() *helper.Response {
	result, err := uc.getAllBank.GetAll()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}
	return helper.GetResponse(result, 200, true)
}
