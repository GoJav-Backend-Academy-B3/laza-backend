package review

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *ReviewRepo) Insert(review model.Review) (cr model.Review, err error) {
	err = r.db.Create(&review).Scan(&cr).Error
	if err != nil {
		return
	}
	return
}
