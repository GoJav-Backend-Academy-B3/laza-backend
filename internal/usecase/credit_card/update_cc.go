package credit_card

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/requests"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type updateCreditCardUsecase struct {
	updateCcRepo         repo.UpdateAction[model.CreditCard]
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction
	validate             *validator.Validate
}

func (uc *updateCreditCardUsecase) Execute(userId, ccId uint64, rb requests.CreditCardRequest) (_result model.CreditCard, statusCode int, err error) {

	err = uc.validate.Struct(rb)
	if err != nil {
		statusCode = 400
		return
	}

	if rb.CVV == "" {
		rb.CVV = "123"
	}
	responseMd, errMd := uc.fetchMidtransCCToken.FetchMidtransCCToken(rb.CardNumber, rb.ExpiredMonth, rb.ExpiredYear, rb.CVV)

	if errMd != nil {
		statusCode = 500
		err = errMd.RawError
		return
	}
	if responseMd.StatusCode != "200" {
		statusCode = 400
		err = errors.New(responseMd.StatusMessage)
		return
	}

	md := model.CreditCard{Id: ccId, UserId: userId, ExpiredMonth: rb.ExpiredMonth, ExpiredYear: rb.ExpiredYear}

	_result, err = uc.updateCcRepo.Update(ccId, md)

	if err != nil {
		statusCode = 500
		return
	}
	return
}

func NewupdateCreditCardUsecase(
	updateCcRepo repo.UpdateAction[model.CreditCard],
	validate *validator.Validate,
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction,
) uc.UpdateCreditCardUsecase {
	return &updateCreditCardUsecase{
		updateCcRepo:         updateCcRepo,
		validate:             validate,
		fetchMidtransCCToken: fetchMidtransCCToken,
	}
}
