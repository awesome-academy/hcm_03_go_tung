package models

import (
	"time"
	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Quantity  int64     `gorm:"type:bigint" json:"quantity"`
	DeletedAt *time.Time `gorm:"type:datetime" json:"deleted_at"`
}
