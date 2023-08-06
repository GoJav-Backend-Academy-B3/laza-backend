package bank

import (
	"github.com/phincon-backend/laza/domain/model"
	action "github.com/phincon-backend/laza/domain/repositories/bank"
	"github.com/phincon-backend/laza/domain/usecases/bank"
	"github.com/phincon-backend/laza/helper"
)

type DeleteBankUsecase struct {
	deleteAction action.DeleteBank[model.Bank]
}

func NewDeleteBankUsecase(repo action.DeleteBank[model.Bank]) bank.DeleteBankUsecase {
	return &DeleteBankUsecase{deleteAction: repo}
}

// Excute implements user.DeleteUserUsecase.
func (uc *DeleteBankUsecase) Execute(id uint64) *helper.Response {
	err := uc.deleteAction.Delete(id)
	if err != nil {
		return helper.GetResponse(err.Error(), 500, true)
	}

	return helper.GetResponse("successfully deleted data Bank", 200, true)
}
