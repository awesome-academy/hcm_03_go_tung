package services

import (
	"HCM_03_GO_TUNG/models"
	"HCM_03_GO_TUNG/repository"
	"errors"
	"strings"
	"time"
	"HCM_03_GO_TUNG/utils"

	"github.com/google/uuid"
	"golang.org/x/text/message"
)

type ProfileService interface {
	GetProfile(userID uuid.UUID) (*models.User, error)
	UpdateProfile(userID uuid.UUID, updates map[string]interface{}) (*models.User, error)
	ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error
}

type profileService struct {
	repo repository.UserRepository
}

func NewProfileService(repo repository.UserRepository) ProfileService {
	return &profileService{repo}
}

func (s *profileService) GetProfile(userID uuid.UUID) (*models.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("invalid user ID"))
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("user not found"))
		}
		return nil, errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("system error: %v", err))
	}

	user.PasswordHash = ""
	return user, nil
}

func (s *profileService) UpdateProfile(userID uuid.UUID, updates map[string]interface{}) (*models.User, error) {
	if userID == uuid.Nil {
		return nil, errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("invalid user ID"))
	}

	allowedFields := map[string]bool{
		"name":       true,
		"first_name": true,
		"last_name":  true,
		"avatar":     true,
		"phone":      true,
		"address":    true,
	}

	filteredUpdates := make(map[string]interface{})
	for key, value := range updates {
		if allowedFields[key] {
			filteredUpdates[key] = value
		}
	}

	if name, ok := filteredUpdates["name"].(string); ok {
		filteredUpdates["name"] = strings.TrimSpace(name)
	}

	filteredUpdates["updated_time"] = time.Now()

	err := s.repo.Update(userID, filteredUpdates)
	if err != nil {
		return nil, errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("failed to update profile"))
	}

	return s.GetProfile(userID)
}

func (s *profileService) ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	if len(newPassword) < 6 {
		return errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("new password must be at least 6 characters"))
	}

	user, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("user not found"))
	}

	if user.Provider == "local" {
		if !utils.CheckPasswordHash(oldPassword, user.PasswordHash) {
			return errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("incorrect old password"))
		}
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New(message.NewPrinter(message.MatchLanguage("en")).Sprintf("failed to hash password"))
	}

	updates := map[string]interface{}{
		"password_hash": hashedPassword,
		"updated_time":  time.Now(),
	}

	return s.repo.Update(userID, updates)
}
