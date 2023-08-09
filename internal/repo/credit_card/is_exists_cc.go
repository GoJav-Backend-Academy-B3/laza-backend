package creditcard

import (
	"github.com/phincon-backend/laza/domain/model"
	"gorm.io/gorm"
)

func (r *CreditCardRepo) IsExistsCc(userId uint64, ccNumber string) (bool, error) {
	err := r.db.Where("user_id = ? and card_number = ?", userId, ccNumber).First(&model.CreditCard{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
