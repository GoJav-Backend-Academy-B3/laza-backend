package credit_card

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
	"github.com/phincon-backend/laza/helper"
)

type AddCreditCardUsecase struct {
	addCcRepo repo.InsertAction[model.CreditCard]
	validate  *validator.Validate
}

func (ad *AddCreditCardUsecase) Execute(userId uint64, rb requests.CreditCardRequest) *helper.Response {

	err := ad.validate.Struct(rb)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusBadRequest, true)
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

func NewaddCreditCardUsecase(addCcRepo repo.InsertAction[model.CreditCard],
	validate *validator.Validate,
) uc.AddCreditCardUsecase {
	return &AddCreditCardUsecase{
		addCcRepo: addCcRepo,
		validate:  validate,
	}
}
