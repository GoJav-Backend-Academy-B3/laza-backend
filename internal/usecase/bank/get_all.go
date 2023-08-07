package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type GetAllBankUsecase struct {
	getAllBank action.GetAllBank[model.Bank]
}

func NewGetAllBankUsecase(repo action.GetAllBank[model.Bank]) bank.GetAllBankUsecase {
	return &GetAllBankUsecase{getAllBank: repo}
}

func (uc *GetAllBankUsecase) Execute() *helper.Response {
	result, err := uc.getAllBank.GetAll()
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}
	return helper.GetResponse(result, 200, true)
}
