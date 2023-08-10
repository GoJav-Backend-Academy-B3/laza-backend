package credit_card

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/phincon-backend/laza/domain/model"
	repo "github.com/phincon-backend/laza/domain/repositories"
	"github.com/phincon-backend/laza/domain/repositories/midtrans"
	"github.com/phincon-backend/laza/domain/requests"
	"github.com/phincon-backend/laza/domain/response"
	uc "github.com/phincon-backend/laza/domain/usecases/credit_card"
	"github.com/phincon-backend/laza/helper"

	"gorm.io/gorm"
)

type updateCreditCardUsecase struct {
	updateCcRepo         repo.UpdateAction[model.CreditCard]
	fetchMidtransCCToken midtrans.FetchMidtransCCTokenAction
	validate             *validator.Validate
}

func (uc *updateCreditCardUsecase) Execute(userId, ccId uint64, rb requests.CreditCardRequest) *helper.Response {

	err := uc.validate.Struct(rb)
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusBadRequest, true)
	}

	if rb.CVV == "" {
		rb.CVV = "123"
	}
	responseMd, errMd := uc.fetchMidtransCCToken.FetchMidtransCCToken(rb.CardNumber, rb.ExpiredMonth, rb.ExpiredYear, rb.CVV)

	if errMd != nil {
		return helper.GetResponse(errMd.RawError.Error(), http.StatusInternalServerError, true)
	}
	if responseMd.StatusCode != "200" {
		return helper.GetResponse(errors.New(responseMd.StatusMessage).Error(), http.StatusBadRequest, true)
	}

	md := model.CreditCard{Id: ccId}
	rb.FilltoField(userId, &md)

	rs, err := uc.updateCcRepo.Update(ccId, md)

	if err == gorm.ErrRecordNotFound {
		return helper.GetResponse(errors.New("credit card not found").Error(), http.StatusNotFound, true)
	}
	if err != nil {
		return helper.GetResponse(err.Error(), http.StatusInternalServerError, true)
	}
	return helper.GetResponse(response.CreditCardResponse{}.FillFromEntity(rs), http.StatusOK, false)
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
