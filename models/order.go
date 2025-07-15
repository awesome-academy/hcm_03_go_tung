package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Status    string    `gorm:"type:enum('pending','completed','cancelled')" json:"status"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
}
