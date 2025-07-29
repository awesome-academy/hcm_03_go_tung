package repository

import (
	"errors"
	"HCM_03_GO_TUNG/models"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// UserRepository định nghĩa interface cho các thao tác DB liên quan đến người dùng
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByProviderID(provider string, providerID string) (*models.User, error)
	Update(id uuid.UUID, updates map[string]interface{}) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create inserts a new user record into the database
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByProviderID(provider string, providerID string) (*models.User, error) {
	var user models.User
	err := r.db.Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error
	return &user, err
}

// Update updates only allowed fields for a user
func (r *userRepository) Update(id uuid.UUID, updates map[string]interface{}) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	// Whitelisted fields that can be updated
	validFields := map[string]struct{}{
		"email":    {},
		"username": {},
	}

	// Validate update fields
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
