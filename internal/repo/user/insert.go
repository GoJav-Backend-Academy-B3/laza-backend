package user

import (
	"github.com/phincon-backend/laza/domain/model"
)

func (r *UserRepo) Insert(dao model.User) (e model.User, err error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return
	}

	daoUser := model.User{
		Username:   dao.Username,
		Password:   dao.Password,
		Email:      dao.Email,
		FullName:   dao.FullName,
		IsVerified: dao.IsVerified,
		IsAdmin:    dao.IsAdmin,
	}

	if err = tx.Create(&daoUser).Scan(&e).Error; err != nil {
		tx.Rollback()
		return
	}

	daoToken := model.VerificationToken{
		Token:      dao.VerificationTokens[0].Token,
		ExpiryDate: dao.VerificationTokens[0].ExpiryDate,
		UserId:     uint64(e.Id),
	}

	if err = tx.Create(&daoToken).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
