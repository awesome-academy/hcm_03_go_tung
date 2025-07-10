package models

import (
	"time"

	"github.com/google/uuid"
)

type ProductReview struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ProductID uuid.UUID `gorm:"type:uuid" json:"product_id"`
	Rating    int64     `gorm:"type:bigint" json:"rating"`
	Comment   string    `gorm:"type:varchar(255)" json:"comment"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
}
