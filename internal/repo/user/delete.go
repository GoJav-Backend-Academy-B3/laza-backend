package user

func (r *UserRepo) Delete(id uint64) error {
	if err := r.db.Delete("id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
