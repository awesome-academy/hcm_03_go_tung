package services

import (
<<<<<<< Updated upstream
	"foods-drinks-app/models"
	"foods-drinks-app/repository"
	"foods-drinks-app/utils"
=======
>>>>>>> Stashed changes
	"errors"
	"foods-drinks-app/models"
	"foods-drinks-app/repository"
	"foods-drinks-app/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ProfileService định nghĩa interface cho các thao tác hồ sơ người dùng
type ProfileService interface {
	GetProfile(userID uuid.UUID) (*models.User, error)
	UpdateProfile(userID uuid.UUID, updates map[string]interface{}) (*models.User, error)
	ChangePassword(userID uuid.UUID, oldPassword, newPassword string, lang utils.Language) error
}

type profileService struct {
	repo repository.UserRepository
}

func NewProfileService(repo repository.UserRepository) ProfileService {
	return &profileService{repo}
}

// GetProfile lấy thông tin hồ sơ người dùng
func (s *profileService) GetProfile(userID uuid.UUID) (*models.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Ẩn thông tin nhạy cảm trước khi trả về
	user.PasswordHash = ""
	return user, nil
}

// UpdateProfile cập nhật thông tin hồ sơ người dùng
func (s *profileService) UpdateProfile(userID uuid.UUID, updates map[string]interface{}) (*models.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID")
	}

	// Danh sách các trường được phép cập nhật
	allowedFields := map[string]bool{
		"name":       true,
		"first_name": true,
		"last_name":  true,
		"avatar":     true,
		"phone":      true,
		"address":    true,
	}

	// Lọc các trường không được phép cập nhật
	filteredUpdates := make(map[string]interface{})
	for key, value := range updates {
		if allowedFields[key] {
			filteredUpdates[key] = value
		}
	}

	// Nếu có cập nhật tên, xử lý chuỗi
	if name, ok := filteredUpdates["name"].(string); ok {
		filteredUpdates["name"] = strings.TrimSpace(name)
	}

	// Thêm thời gian cập nhật
	filteredUpdates["updated_time"] = time.Now()

	err := s.repo.Update(userID, filteredUpdates)
	if err != nil {
		return nil, errors.New("failed to update profile")
	}

	// Lấy lại thông tin user sau khi cập nhật
	return s.GetProfile(userID)
}

// ChangePassword thay đổi mật khẩu người dùng
func (s *profileService) ChangePassword(userID uuid.UUID, oldPassword, newPassword string, lang utils.Language) error {
	if len(newPassword) < 6 {
		return errors.New(utils.GetMessage(lang, utils.ErrorInvalidPassword))
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New(utils.GetMessage(lang, utils.ErrorUserNotFound))
	}

	// Kiểm tra mật khẩu cũ (nếu là tài khoản local)
	if user.Provider == "local" {
		if !utils.CheckPasswordHash(oldPassword, user.PasswordHash) {
			return errors.New(utils.GetMessage(lang, utils.ErrorIncorrectOldPassword))
		}
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(utils.GetMessage(lang, utils.ErrorFailedToHashPassword))
	}

	updates := map[string]interface{}{
		"password_hash": hashedPassword,
		"updated_time":  time.Now(),
	}

	return s.repo.Update(userID, updates)
}
