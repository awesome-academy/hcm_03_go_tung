package models

import (
	"time"

	"github.com/google/uuid"
)

// Order status constants
const (
	OrderStatusPending   = "pending"
	OrderStatusCompleted = "completed"
	OrderStatusCancelled = "cancelled"
)

type Order struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID          uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Status          string    `gorm:"type:varchar(20)" json:"status"`
	TotalAmount     float64   `gorm:"type:decimal(10,2)" json:"total_amount"`
	CustomerName    string    `gorm:"type:varchar(255)" json:"customer_name"`
	CustomerPhone   string    `gorm:"type:varchar(20)" json:"customer_phone"`
	DeliveryAddress string    `gorm:"type:text" json:"delivery_address"`
	CreatedAt       time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:datetime" json:"updated_at"`
}
