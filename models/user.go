package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Email       string    `gorm:"type:varchar(255);unique" json:"email"`
	Role        string    `gorm:"type:enum('admin','user')" json:"role"`
	CreatedTime time.Time `gorm:"type:timestamp" json:"created_time"`
}
