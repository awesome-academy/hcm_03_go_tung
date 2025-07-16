package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductSuggestion struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
}
