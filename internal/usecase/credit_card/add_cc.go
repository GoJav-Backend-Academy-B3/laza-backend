package credit_card

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	rp "github.com/phincon-backend/laza/domain/repositories/credit_card"
	"github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/requests"

	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
)

type AddCreditCardUsecase struct {
	isExistsCc           rp.IsExistsCcAction
	addCcRepo            repo.InsertAction[model.CreditCard]
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction
	validate             *validator.Validate
}

func (ad *AddCreditCardUsecase) Execute(userId uint64, rb requests.CreditCardRequest) (_result model.CreditCard, statusCode int, err error) {

	err = ad.validate.Struct(rb)
	if err != nil {
		statusCode = 400
		return
	}

	if rb.CVV == "" {
		rb.CVV = "123"
	}
	responseMd, errMd := ad.fetchMidtransCCToken.FetchMidtransCCToken(rb.CardNumber, rb.ExpiredMonth, rb.ExpiredYear, rb.CVV)

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

	tf, err := ad.isExistsCc.IsExistsCc(userId, rb.CardNumber)
	if err != nil {
		statusCode = 500
		return
	}
	if tf {
		err = errors.New("credit card already saved")
		statusCode = 400
		return
	}

	md := model.CreditCard{UserId: userId, CardNumber: rb.CardNumber, ExpiredMonth: rb.ExpiredMonth, ExpiredYear: rb.ExpiredYear}

	_result, err = ad.addCcRepo.Insert(md)

	return
}

func NewaddCreditCardUsecase(
	isExistsCc rp.IsExistsCcAction,
	addCcRepo repo.InsertAction[model.CreditCard],
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction,
	validate *validator.Validate,
) uc.AddCreditCardUsecase {

	return &AddCreditCardUsecase{
		isExistsCc:           isExistsCc,
		addCcRepo:            addCcRepo,
		fetchMidtransCCToken: fetchMidtransCCToken,
		validate:             validate,
	}
}
