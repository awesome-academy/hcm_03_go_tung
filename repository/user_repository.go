package repository

import (
	"gorm.io/gorm"
	"HCM_03_GO_TUNG/models"
)

// UserRepository định nghĩa interface cho các thao tác DB liên quan đến người dùng
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByProviderID(provider string, providerID string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create inserts a new user record into the database
// Returns error if the operation fails
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string, preloads ...string) (*models.User, error) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByProviderID(provider string, providerID string) (*models.User, error) {
	var user models.User

	err := r.db.Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error

	return &user, err
}

