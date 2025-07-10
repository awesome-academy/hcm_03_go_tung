package services

import (
	"errors"
	"fmt"
	"foods-drinks-app/models"
	"foods-drinks-app/repository"
	"foods-drinks-app/utils"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserService định nghĩa interface cho logic xử lý liên quan đến người dùng
type UserService interface {
	SignUp(email, password, firstName, lastName string) (*models.User, error)
	Login(email, password string) (*models.User, error)
	LoginWithGoogle(provider, providerID, name, email string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

// SignUp đăng ký người dùng mới từ thông tin email/password
func (s *userService) SignUp(email, password, firstName, lastName string) (*models.User, error) {
	// Validate email and password
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Kiểm tra email đã tồn tại chưa
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil && existing.ID != uuid.Nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:           uuid.New(),
		Name:         strings.TrimSpace(firstName + " " + lastName),
		Email:        email,
		PasswordHash: hashedPassword,
		Provider:     "local",
		Role:         "user",
		IsActive:     true,
		CreatedTime:  time.Now(),
	}
	return user, s.repo.Create(user)
}

// Login xử lý đăng nhập bằng email/password
func (s *userService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user.ID == uuid.Nil {
		return nil, errors.New("invalid credentials")
	}
	// IsActive để đảm bảo user chưa bị khóa hoặc vô hiệu hóa
	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// Login With Google xử lý đăng nhập bằng OAuth Google
func (s *userService) LoginWithGoogle(provider, providerID, name, email string) (*models.User, error) {
	user, err := s.repo.FindByProviderID(provider, providerID)

	// Nếu tìm thấy user → trả về luôn
	if err == nil && user != nil && user.ID != uuid.Nil {
		return user, nil
	}

	// Nếu lỗi KHÔNG PHẢI là "record not found" → đây là lỗi hệ thống
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("failed to query user by provider ID: %w", err)
	}

	// Nếu là record not found → tiếp tục tạo user mới
	newUser := &models.User{
		ID:          uuid.New(),
		Name:        name,
		Email:       email,
		Provider:    provider,
		ProviderID:  providerID,
		Role:        "user",
		IsActive:    true,
		CreatedTime: time.Now(),
	}
	return newUser, s.repo.Create(newUser)
}
