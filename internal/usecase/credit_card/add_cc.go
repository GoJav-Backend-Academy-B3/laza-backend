package credit_card

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	rp "github.com/phincon-backend/laza/domain/repositories/credit_card"

	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
	"github.com/phincon-backend/laza/helper"
)

type addCreditCardUsecase struct {
	isExistsCc rp.IsExistsCcAction
	addCcRepo  repo.InsertAction[model.CreditCard]
	validate   *validator.Validate
}

func (ad *addCreditCardUsecase) Execute(userId uint64, rb requests.CreditCardRequest) *helper.Response {

	err := ad.validate.Struct(rb)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusBadRequest, true)
	}

	tf, err := ad.isExistsCc.IsExistsCc(userId, rb.CardNumber)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	if tf {
		return helper.GetResponse("credit card already saved", http.StatusInternalServerError, true)
	}

	md := model.CreditCard{}
	rb.FilltoField(userId, &md)

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
	validate *validator.Validate,
) uc.AddCreditCardUsecase {
	return &addCreditCardUsecase{
		isExistsCc: isExistsCc,
		addCcRepo:  addCcRepo,
		validate:   validate,
	}
}
