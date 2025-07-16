package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	Email        string    `gorm:"type:varchar(255);unique" json:"email"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"-"` // Lưu mật khẩu đã mã hóa
	AvatarURL    string    `gorm:"type:varchar(255)" json:"avatar_url"`
	Provider     string    `gorm:"type:varchar(50)" json:"provider"`     // local, facebook, google
	ProviderID   string    `gorm:"type:varchar(255)" json:"provider_id"` // ID từ OAuth
	Role         string    `gorm:"type:enum('admin','user')" json:"role"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	LastLogin    time.Time `gorm:"type:timestamp" json:"last_login"`
	Phone        string    `gorm:"type:varchar(20)" json:"phone"`
	Address      string    `gorm:"type:varchar(255)" json:"address"`
	CreatedTime  time.Time `gorm:"type:timestamp" json:"created_time"`
}
