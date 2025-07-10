package repository

import (
	"errors"
	"foods-drinks-app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository định nghĩa interface cho các thao tác DB liên quan đến người dùng
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uuid.UUID) (*models.User, error)
	FindByEmail(email string, preloads ...string) (*models.User, error)
	FindByProviderID(provider string, providerID string) (*models.User, error)
	Update(id uuid.UUID, updates map[string]interface{}) error
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create thêm người dùng mới vào database
func (r *userRepository) Create(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}
	return r.db.Create(user).Error
}

// FindByID tìm người dùng bằng ID
func (r *userRepository) FindByID(id uuid.UUID) (*models.User, error) {
	if id == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail tìm người dùng bằng email với khả năng preload
func (r *userRepository) FindByEmail(email string, preloads ...string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	query := r.db.Where("email = ?", email)

	// Thêm preload nếu có
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	var user models.User
	err := query.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByProviderID tìm người dùng bằng provider ID
func (r *userRepository) FindByProviderID(provider string, providerID string) (*models.User, error) {
	if provider == "" || providerID == "" {
		return nil, errors.New("provider and provider ID cannot be empty")
	}

	var user models.User
	err := r.db.Where("provider = ? AND provider_id = ?", provider, providerID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update cập nhật thông tin người dùng
func (r *userRepository) Update(id uuid.UUID, updates map[string]interface{}) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
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

// Delete xóa người dùng (soft delete)
func (r *userRepository) Delete(id uuid.UUID) error {
	if id == uuid.Nil {
		return errors.New("invalid user ID")
	}

	result := r.db.Delete(&models.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}
