package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Type      string    `gorm:"type:char(50)" json:"type"`
	Category  string    `gorm:"type:char(50)" json:"category"`
	Price     float64   `gorm:"type:float" json:"price"`
	ImageURL  string    `gorm:"type:varchar(255)" json:"image_url"`
	Rating    float64   `gorm:"type:float" json:"rating"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
