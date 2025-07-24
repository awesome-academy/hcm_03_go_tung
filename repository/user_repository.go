
func (r *userRepository) Update(id uuid.UUID, updates map[string]interface{}) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	validFields := map[string]struct{}{
		"email":    {},
		"username": {},

	}


	for field := range updates {
		if _, ok := validFields[field]; !ok {
			return errors.New("invalid field: " + field)
		}
	}

	result := r.db.Model(&models.User{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
