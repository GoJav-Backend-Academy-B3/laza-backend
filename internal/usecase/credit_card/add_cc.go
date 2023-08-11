package credit_card

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	rp "github.com/phincon-backend/laza/domain/repositories/credit_card"
	"github.com/phincon-backend/laza/domain/repositories/midtrans"

	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
	"github.com/phincon-backend/laza/helper"
)

type AddCreditCardUsecase struct {
	isExistsCc           rp.IsExistsCcAction
	addCcRepo            repo.InsertAction[model.CreditCard]
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction
	validate             *validator.Validate
}

func (ad *AddCreditCardUsecase) Execute(userId uint64, rb requests.CreditCardRequest) *helper.Response {

	err := ad.validate.Struct(rb)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusBadRequest, true)
	}

	if rb.CVV == "" {
		rb.CVV = "123"
	}
	responseMd, errMd := ad.fetchMidtransCCToken.FetchMidtransCCToken(rb.CardNumber, rb.ExpiredMonth, rb.ExpiredYear, rb.CVV)

	if errMd != nil {
		return helper.GetResponse(errMd.RawError.Error(), http.StatusInternalServerError, true)
	}
	if responseMd.StatusCode != "200" {
		return helper.GetResponse(errors.New(responseMd.StatusMessage).Error(), http.StatusBadRequest, true)
	}

	tf, err := ad.isExistsCc.IsExistsCc(userId, rb.CardNumber)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	if tf {
		return helper.GetResponse("credit card already saved", http.StatusBadRequest, true)
	}

	md := model.CreditCard{UserId: userId, CardNumber: rb.CardNumber, ExpiredMonth: rb.ExpiredMonth, ExpiredYear: rb.ExpiredYear}

	rs, err := ad.addCcRepo.Insert(md)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}

	data := response.CreditCardResponse{}.FillFromEntity(rs)
	return helper.GetResponse(data, http.StatusCreated, false)
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
