package cart

import (
	"errors"

	"github.com/phincon-backend/laza/domain/model"
)

func (r *CartRepo) Delete(id any) (err error) {
	m, ok := id.(map[string]uint64)
	if !ok {
		err = errors.New("there is an error in the delete cart repo")
		return
	}

	err = r.db.Where("user_id = ? AND product_id = ?", m["userId"], m["productId"]).Delete(&model.Cart{}).Error
	return
}
